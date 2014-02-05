package cfitsio

import (
   "testing"
)

func TestRicecomp(t *testing.T) {
   input := make([]int32, 4096)

   for i := range input {
      input[i] = int32(i-2048)
   }
   t.Log("input",input[:10])

   buf := Fits_rcomp(input, 32)
   if buf == nil {
      t.Fatal("Rice compression failed.")
   }
   t.Log("buf",buf[:10])

   output := Fits_rdecomp(buf, 4096, 32)
   if output == nil {
      t.Fatal("Rice decompression failed.")
   }

   t.Log("output",output[:10])
   var cnt int = 0
   for i := range input {
      if input[i]  != output[i] {
         cnt++
      }
   }
   if cnt > 0 {
      t.Fatalf("%v values not equal.\n", cnt)
   }
}

func TestRicecompByte(t *testing.T) {
   input := make([]int8, 4096)

   for i := range input {
      input[i] = int8((i%256)-127)
   }
   t.Log("input",input[:10])

   buf := Fits_rcomp_byte(input, 32)
   if buf == nil {
      t.Fatal("Rice compression failed.")
   }
   t.Log("buf",buf[:10])

   output := Fits_rdecomp_byte(buf, 4096, 32)
   if output == nil {
      t.Fatal("Rice decompression failed.")
   }

   t.Log("output",output[:10])
   var cnt int = 0
   for i := range input {
      if input[i]  != output[i] {
         cnt++
      }
   }
   if cnt > 0 {
      t.Fatalf("%v values not equal.\n", cnt)
   }
}

func TestRicecompShort(t *testing.T) {
   input := make([]int16, 4096)

   for i := range input {
      input[i] = int16(i-2048)
   }
   t.Log("input",input[:10])

   buf := Fits_rcomp_short(input, 32)
   if buf == nil {
      t.Fatal("Rice compression failed.")
   }
   t.Log("buf",buf[:10])

   output := Fits_rdecomp_short(buf, 4096, 32)
   if output == nil {
      t.Fatal("Rice decompression failed.")
   }

   t.Log("output",output[:10])
   var cnt int = 0
   for i := range input {
      if input[i]  != output[i] {
         cnt++
      }
   }
   if cnt > 0 {
      t.Fatalf("%v values not equal.\n", cnt)
   }
}
