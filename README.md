# []byte to other conversion benchmarks

## Jun 16, 2024

Ran on a Haswell-based Xeon.

```
cpu: Intel Core Processor (Haswell, no TSX, IBRS)
BenchmarkUint32_StdlibReader-48              	30774943	        54.22 ns/op	       4 B/op	       1 allocs/op
BenchmarkUint32_StdlibReflect-48             	23037266	        49.82 ns/op	       4 B/op	       1 allocs/op
BenchmarkUint32_SliceDirect-48               	807744319	         1.600 ns/op	       0 B/op	       0 allocs/op
BenchmarkUint32_BufferNext-48                	277967260	         5.064 ns/op	       0 B/op	       0 allocs/op
BenchmarkUint32_BufferNextChecked-48         	233251548	         7.312 ns/op	       0 B/op	       0 allocs/op
BenchmarkUint32_BufferReadFull-48            	29868187	        37.25 ns/op	       4 B/op	       1 allocs/op
BenchmarkUint32_BufferRead-48                	130718455	         7.658 ns/op	       0 B/op	       0 allocs/op
BenchmarkUint32_CustomReader-48              	506034272	         2.017 ns/op	       0 B/op	       0 allocs/op
BenchmarkUint32_DynReadBuffer-48             	70568961	        23.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkUint32_DynReadReader-48             	23001098	        44.27 ns/op	       4 B/op	       1 allocs/op
BenchmarkByte_ReadFull-48                    	32636216	        32.21 ns/op	       1 B/op	       1 allocs/op
BenchmarkByte_ByteReader-48                  	301626748	         3.818 ns/op	       0 B/op	       0 allocs/op
BenchmarkByte_DynByteReader-48               	70775143	        19.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkBytes_ReaderReadFull-48             	80370420	        21.93 ns/op	       0 B/op	       0 allocs/op
BenchmarkBytes_ReaderCopyN-48                	 9804448	       129.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkBytes_ReaderCopyBuffer-48           	 9269085	       116.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkBytes_BufferCopyN-48                	 7110308	       173.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkBytes_BufferDynCopy-48              	58197999	        18.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkString_ReadStringify8-48            	26089701	        47.18 ns/op	       8 B/op	       1 allocs/op
BenchmarkString_ReadStringify64-48           	20329069	       106.7 ns/op	      64 B/op	       1 allocs/op
BenchmarkString_ReadStringify256-48          	11340177	       515.8 ns/op	     256 B/op	       1 allocs/op
BenchmarkString_ReadStringify1024-48         	 2945204	      1400 ns/op	    1024 B/op	       1 allocs/op
BenchmarkString_ReadStringify4096-48         	  650972	      8360 ns/op	    4096 B/op	       1 allocs/op
BenchmarkString_ReadStringify16384-48        	  150310	     18102 ns/op	   16384 B/op	       1 allocs/op
BenchmarkString_CopyBufferBuilder8-48        	 6616008	       221.2 ns/op	      40 B/op	       2 allocs/op
BenchmarkString_CopyBufferBuilder64-48       	 4669700	       222.2 ns/op	      96 B/op	       2 allocs/op
BenchmarkString_CopyBufferBuilder256-48      	 2521032	       525.1 ns/op	     288 B/op	       2 allocs/op
BenchmarkString_CopyBufferBuilder1024-48     	 1000000	      1300 ns/op	    1056 B/op	       2 allocs/op
BenchmarkString_CopyBufferBuilder4096-48     	  500678	      4080 ns/op	    4128 B/op	       2 allocs/op
BenchmarkString_CopyBufferBuilder16384-48    	  163058	     13596 ns/op	   16416 B/op	       2 allocs/op
BenchmarkString_CopyReuseBuilder8-48         	17296796	        67.52 ns/op	       8 B/op	       1 allocs/op
BenchmarkString_CopyReuseBuilder64-48        	 9499611	       148.3 ns/op	      64 B/op	       1 allocs/op
BenchmarkString_CopyReuseBuilder256-48       	 2119345	       544.7 ns/op	     256 B/op	       1 allocs/op
BenchmarkString_CopyReuseBuilder1024-48      	 1000000	      1510 ns/op	    1024 B/op	       1 allocs/op
BenchmarkString_CopyReuseBuilder4096-48      	  396292	      5239 ns/op	    4096 B/op	       1 allocs/op
BenchmarkString_CopyReuseBuilder16384-48     	  155736	     14759 ns/op	   16384 B/op	       1 allocs/op
```
