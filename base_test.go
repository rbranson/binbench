package binbench

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"strings"
	"testing"
)

var junku32 uint32
var junkbyte byte
var junkbytes []byte
var junkbuf *bytes.Buffer
var junkstring string

func uint32Fill(n int) []byte {
	buf := make([]byte, 0, 4*n)
	for i := 0; i < n; i++ {
		buf = binary.LittleEndian.AppendUint32(buf, uint32(i))
	}
	return buf
}

func byteFill(n int) []byte {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = byte(i)
	}
	return buf
}

func readUint32StdlibReader(r io.Reader) (val uint32, err error) {
	buf := make([]byte, 4)
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return
	}
	val = binary.LittleEndian.Uint32(buf)
	return
}

func BenchmarkUint32_StdlibReader(b *testing.B) {
	b.ReportAllocs()
	bb := bytes.NewReader(uint32Fill(b.N))
	b.ResetTimer()
	var total uint32
	var err error
	for n := 0; n < b.N; n++ {
		total, err = readUint32StdlibReader(bb)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
	}
	junku32 = total
}

func readUint32StdlibReflect(r io.Reader) (val uint32, err error) {
	err = binary.Read(r, binary.LittleEndian, &val)
	return
}

func BenchmarkUint32_StdlibReflect(b *testing.B) {
	b.ReportAllocs()
	bb := bytes.NewReader(uint32Fill(b.N))
	b.ResetTimer()
	var total uint32
	var err error
	for n := 0; n < b.N; n++ {
		total, err = readUint32StdlibReflect(bb)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
	}
	junku32 = total
}

func readUint32FromSlice(b []byte) (ob []byte, val uint32) {
	val = binary.LittleEndian.Uint32(b)
	ob = b[4:]
	return
}

func BenchmarkUint32_SliceDirect(b *testing.B) {
	b.ReportAllocs()
	buf := uint32Fill(b.N)
	b.ResetTimer()
	var total uint32
	for n := 0; n < b.N; n++ {
		buf, total = readUint32FromSlice(buf)
	}
	junku32 = total
}

func readUint32BufferNext(buf *bytes.Buffer) uint32 {
	return binary.LittleEndian.Uint32(buf.Next(4))
}

func BenchmarkUint32_BufferNext(b *testing.B) {
	b.ReportAllocs()
	bb := bytes.NewBuffer(uint32Fill(b.N))
	b.ResetTimer()
	var total uint32
	for n := 0; n < b.N; n++ {
		total = readUint32BufferNext(bb)
	}
	junku32 = total
}

func readUint32BufferNextChecked(buf *bytes.Buffer) (val uint32, err error) {
	if buf.Len() < 4 {
		err = errors.New("not enough buffer")
		return
	}
	return binary.LittleEndian.Uint32(buf.Next(4)), nil
}

func BenchmarkUint32_BufferNextChecked(b *testing.B) {
	b.ReportAllocs()
	bb := bytes.NewBuffer(uint32Fill(b.N))
	b.ResetTimer()
	var total uint32
	var err error
	for n := 0; n < b.N; n++ {
		total, err = readUint32BufferNextChecked(bb)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
	}
	junku32 = total
}

func readUint32BufferReadFull(buf *bytes.Buffer) (val uint32, err error) {
	var b [4]byte
	_, err = io.ReadFull(buf, b[:])
	val = binary.LittleEndian.Uint32(b[:])
	return
}

func BenchmarkUint32_BufferReadFull(b *testing.B) {
	b.ReportAllocs()
	bb := bytes.NewBuffer(uint32Fill(b.N))
	b.ResetTimer()
	var total uint32
	var err error
	for n := 0; n < b.N; n++ {
		total, err = readUint32BufferReadFull(bb)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
	}
	junku32 = total
}

func readUint32BufferRead(buf *bytes.Buffer) (val uint32, err error) {
	var b [4]byte
	_, err = buf.Read(b[:])
	val = binary.LittleEndian.Uint32(b[:])
	return
}

func BenchmarkUint32_BufferRead(b *testing.B) {
	b.ReportAllocs()
	bb := bytes.NewBuffer(uint32Fill(b.N))
	b.ResetTimer()
	var total uint32
	var err error
	for n := 0; n < b.N; n++ {
		total, err = readUint32BufferRead(bb)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
	}
	junku32 = total
}

type customReader struct {
	buf []byte
	off int
}

func (r *customReader) ReadUint32() (val uint32, err error) {
	if len(r.buf)-r.off < 4 {
		err = errors.New("not enough buffer")
		return
	}
	val = binary.LittleEndian.Uint32(r.buf[r.off:])
	r.off += 4
	return
}

func BenchmarkUint32_CustomReader(b *testing.B) {
	b.ReportAllocs()
	rd := customReader{buf: uint32Fill(b.N)}
	b.ResetTimer()
	var total uint32
	var err error
	for n := 0; n < b.N; n++ {
		total, err = rd.ReadUint32()
		if err != nil {
			b.Fatalf("error: %v", err)
		}
	}
	junku32 = total
}

type next interface {
	Len() int
	Next(n int) []byte
}

func readUint32BDynReadNext(r io.Reader) (val uint32, err error) {
	if n, ok := r.(next); ok {
		if n.Len() < 4 {
			err = errors.New("not enough buffer")
			return
		}
		return binary.LittleEndian.Uint32(n.Next(4)), nil
	}
	var b [4]byte
	_, err = io.ReadFull(r, b[:])
	val = binary.LittleEndian.Uint32(b[:])
	return
}

func BenchmarkUint32_DynReadBuffer(b *testing.B) {
	b.ReportAllocs()
	bb := bytes.NewBuffer(uint32Fill(b.N))
	b.ResetTimer()
	var total uint32
	var err error
	for n := 0; n < b.N; n++ {
		total, err = readUint32BDynReadNext(bb)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
	}
	junku32 = total
}

func BenchmarkUint32_DynReadReader(b *testing.B) {
	b.ReportAllocs()
	bb := bytes.NewReader(uint32Fill(b.N))
	b.ResetTimer()
	var total uint32
	var err error
	for n := 0; n < b.N; n++ {
		total, err = readUint32BDynReadNext(bb)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
	}
	junku32 = total
}

func readByteReadFull(r io.Reader) (val byte, err error) {
	var buf [1]byte
	_, err = io.ReadFull(r, buf[:])
	if err != nil {
		return
	}
	return buf[0], nil
}

func BenchmarkByte_ReadFull(b *testing.B) {
	b.ReportAllocs()
	rd := bytes.NewReader(byteFill(b.N))
	b.ResetTimer()
	var total byte
	var err error
	for n := 0; n < b.N; n++ {
		total, err = readByteReadFull(rd)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
	}
	junkbyte = total
}

func readByteByteReader(r io.ByteReader) (val byte, err error) {
	return r.ReadByte()
}

func BenchmarkByte_ByteReader(b *testing.B) {
	b.ReportAllocs()
	rd := bytes.NewReader(byteFill(b.N))
	b.ResetTimer()
	var total byte
	var err error
	for n := 0; n < b.N; n++ {
		total, err = readByteByteReader(rd)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
	}
	junkbyte = total
}

func readByteDyn(r io.Reader) (byte, error) {
	if br, ok := r.(io.ByteReader); ok {
		return readByteByteReader(br)
	}
	return readByteReadFull(r)
}

func BenchmarkByte_DynByteReader(b *testing.B) {
	b.ReportAllocs()
	rd := bytes.NewReader(byteFill(b.N))
	b.ResetTimer()
	var total byte
	var err error
	for n := 0; n < b.N; n++ {
		total, err = readByteDyn(rd)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
	}
	junkbyte = total
}

func readBytesReaderReadFull(r io.Reader, buf []byte) ([]byte, error) {
	_, err := io.ReadFull(r, buf)
	return buf, err
}

func BenchmarkBytes_ReaderReadFull(b *testing.B) {
	b.ReportAllocs()
	rd := bytes.NewReader(byteFill(b.N * 8))
	b.ResetTimer()
	retain := make([]byte, 8)
	var err error
	for n := 0; n < b.N; n++ {
		retain, err = readBytesReaderReadFull(rd, retain)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
	}
	junkbytes = retain
}

func readBytesReaderCopyN(r io.Reader, w io.Writer, n int) error {
	_, err := io.CopyN(w, r, int64(n))
	return err
}

func BenchmarkBytes_ReaderCopyN(b *testing.B) {
	b.ReportAllocs()
	rd := bytes.NewReader(byteFill(b.N * 8))
	b.ResetTimer()
	var retain bytes.Buffer
	var err error
	for n := 0; n < b.N; n++ {
		err = readBytesReaderCopyN(rd, &retain, 8)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
		retain.Reset()
	}
}

func readBytesReaderCopyBuffer(r io.Reader, w io.Writer, n int, buf []byte) error {
	_, err := io.CopyBuffer(w, io.LimitReader(r, int64(n)), buf)
	return err
}

func BenchmarkBytes_ReaderCopyBuffer(b *testing.B) {
	b.ReportAllocs()
	rd := bytes.NewReader(byteFill(b.N * 8))
	b.ResetTimer()
	var retain bytes.Buffer
	var err error
	tmp := make([]byte, 512)
	for n := 0; n < b.N; n++ {
		err = readBytesReaderCopyBuffer(rd, &retain, 8, tmp)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
		retain.Reset()
	}
	junkbuf = &retain
}

func readBytesBufferCopyN(r io.Reader, w io.Writer, n int) error {
	_, err := io.CopyN(w, r, int64(n))
	return err
}

func BenchmarkBytes_BufferCopyN(b *testing.B) {
	b.ReportAllocs()
	rd := bytes.NewBuffer(byteFill(b.N * 8))
	b.ResetTimer()
	var retain bytes.Buffer
	var err error
	for n := 0; n < b.N; n++ {
		err = readBytesBufferCopyN(rd, &retain, 8)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
		retain.Reset()
	}
	junkbuf = &retain
}

func readBytesBufferDynCopy(r io.Reader, w io.Writer, n int) error {
	if rb, ok := r.(*bytes.Buffer); ok {
		if wb, ok := w.(*bytes.Buffer); ok {
			if rb.Len() < n {
				return errors.New("not enough buffer")
			}
			buf := wb.AvailableBuffer()
			buf = append(buf, rb.Next(n)...)
			_, err := wb.Write(buf)
			return err
		}
	}

	_, err := io.CopyN(w, r, int64(n))
	return err
}

func BenchmarkBytes_BufferDynCopy(b *testing.B) {
	b.ReportAllocs()
	bf := byteFill(8 * b.N)
	rd := bytes.NewBuffer(bf)
	wb := bytes.NewBuffer(make([]byte, 0, 8))
	b.ResetTimer()
	var err error
	for n := 0; n < b.N; n++ {
		err = readBytesBufferDynCopy(rd, wb, 8)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
		wb.Reset()
	}
}

func benchmarkString_ReadStringify(size int, b *testing.B) {
	b.ReportAllocs()
	bf := byteFill(size * b.N)
	rd := bytes.NewBuffer(bf)
	b.ResetTimer()
	var err error
	var s string
	buf := make([]byte, size)
	for n := 0; n < b.N; n++ {
		_, err = io.ReadFull(rd, buf)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
		s = string(buf)
	}
	junkstring = s
}

func BenchmarkString_ReadStringify8(b *testing.B)     { benchmarkString_ReadStringify(8, b) }
func BenchmarkString_ReadStringify64(b *testing.B)    { benchmarkString_ReadStringify(64, b) }
func BenchmarkString_ReadStringify256(b *testing.B)   { benchmarkString_ReadStringify(256, b) }
func BenchmarkString_ReadStringify1024(b *testing.B)  { benchmarkString_ReadStringify(1024, b) }
func BenchmarkString_ReadStringify4096(b *testing.B)  { benchmarkString_ReadStringify(4096, b) }
func BenchmarkString_ReadStringify16384(b *testing.B) { benchmarkString_ReadStringify(16384, b) }

func benchmarkString_CopyBufferBuilder(size int, b *testing.B) {
	b.ReportAllocs()
	bf := byteFill(size)
	buffers := make([]*bytes.Buffer, b.N)
	for n := 0; n < b.N; n++ {
		buf := append([]byte(nil), bf...)
		buffers[n] = bytes.NewBuffer(buf)
	}

	b.ResetTimer()
	var err error
	var s string
	buf := make([]byte, 4096)
	for n := 0; n < b.N; n++ {
		var sb strings.Builder
		_, err = io.CopyBuffer(&sb, buffers[n], buf)
		if err != nil {
			b.Fatalf("error: %v", err)
		}
		s = sb.String()
	}
	junkstring = s
}

func BenchmarkString_CopyBufferBuilder8(b *testing.B)    { benchmarkString_CopyBufferBuilder(8, b) }
func BenchmarkString_CopyBufferBuilder64(b *testing.B)   { benchmarkString_CopyBufferBuilder(64, b) }
func BenchmarkString_CopyBufferBuilder256(b *testing.B)  { benchmarkString_CopyBufferBuilder(256, b) }
func BenchmarkString_CopyBufferBuilder1024(b *testing.B) { benchmarkString_CopyBufferBuilder(1024, b) }
func BenchmarkString_CopyBufferBuilder4096(b *testing.B) { benchmarkString_CopyBufferBuilder(4096, b) }
func BenchmarkString_CopyBufferBuilder16384(b *testing.B) {
	benchmarkString_CopyBufferBuilder(16384, b)
}

func benchmarkString_CopyReuseBuilder(size int, b *testing.B) {
	b.ReportAllocs()

	bf := byteFill(size)
	buffers := make([]*bytes.Buffer, b.N)
	for n := 0; n < b.N; n++ {
		buf := append([]byte(nil), bf...)
		buffers[n] = bytes.NewBuffer(buf)
	}

	b.ResetTimer()

	var err error
	var s string
	var sb strings.Builder
	for n := 0; n < b.N; n++ {
		sb.Reset()
		_, err = io.Copy(&sb, buffers[n])
		if err != nil {
			b.Fatalf("error: %v", err)
		}
		s = sb.String()
	}
	junkstring = s
}

func BenchmarkString_CopyReuseBuilder8(b *testing.B)     { benchmarkString_CopyReuseBuilder(8, b) }
func BenchmarkString_CopyReuseBuilder64(b *testing.B)    { benchmarkString_CopyReuseBuilder(64, b) }
func BenchmarkString_CopyReuseBuilder256(b *testing.B)   { benchmarkString_CopyReuseBuilder(256, b) }
func BenchmarkString_CopyReuseBuilder1024(b *testing.B)  { benchmarkString_CopyReuseBuilder(1024, b) }
func BenchmarkString_CopyReuseBuilder4096(b *testing.B)  { benchmarkString_CopyReuseBuilder(4096, b) }
func BenchmarkString_CopyReuseBuilder16384(b *testing.B) { benchmarkString_CopyReuseBuilder(16384, b) }
