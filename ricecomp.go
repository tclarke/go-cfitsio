package cfitsio

// #include "go-cfitsio.h"
// #include "go-cfitsio-utils.h"
/*
extern int fits_rcomp(int a[], int nx, unsigned char *c, int clen, int nblock);
extern int fits_rcomp_byte(signed char a[], int nx, unsigned char *c, int clen, int nblock);
extern int fits_rcomp_short(short a[], int nx, unsigned char *c, int clen, int nblock);
extern int fits_rdecomp(unsigned char *c, int clen, int array[], int nx, int nblock);
extern int fits_rdecomp_byte(unsigned char *c, int clen, char array[], int nx, int nblock);
extern int fits_rdecomp_short(unsigned char *c, int clen, short array[], int nx, int nblock);
#include <string.h>
*/
import "C"
import (
   "unsafe"
)

func Fits_rcomp(a []int32, nblock int) []uint8 {
   nx := C.int(len(a))
   buf := make([]uint8, nx*4)

   c_a := (*C.int)(unsafe.Pointer(&a[0]))
   c_buf := (*C.uchar)(unsafe.Pointer(&buf[0]))
   buflen := int(C.fits_rcomp(c_a, nx, c_buf, nx*4, C.int(nblock)))
   if buflen <= 0 {
      return nil
   }
   return buf
}

func Fits_rcomp_byte(a []int8, nblock int) []uint8 {
   nx := C.int(len(a))
   buf := make([]uint8, nx)

   c_a := (*C.schar)(unsafe.Pointer(&a[0]))
   c_buf := (*C.uchar)(unsafe.Pointer(&buf[0]))
   buflen := int(C.fits_rcomp_byte(c_a, nx, c_buf, nx, C.int(nblock)))
   if buflen <= 0 {
      return nil
   }
   return buf
}

func Fits_rcomp_short(a []int16, nblock int) []uint8 {
   nx := C.int(len(a))
   buf := make([]uint8, nx*2)

   c_a := (*C.short)(unsafe.Pointer(&a[0]))
   c_buf := (*C.uchar)(unsafe.Pointer(&buf[0]))
   buflen := int(C.fits_rcomp_short(c_a, nx, c_buf, nx*2, C.int(nblock)))
   if buflen <= 0 {
      return nil
   }
   return buf
}

func Fits_rdecomp(buf []uint8, nx int, nblock int) []int32 {
   clen := (C.int)(len(buf))
   a := make([]int32, nx)

   c_a := (*C.int)(unsafe.Pointer(&a[0]))
   c_buf := (*C.uchar)(unsafe.Pointer(&buf[0]))
   if int(C.fits_rdecomp(c_buf, clen, c_a, C.int(nx), C.int(nblock))) != 0 {
      return nil
   }
   return a
}

func Fits_rdecomp_byte(buf []uint8, nx int, nblock int) []int8 {
   clen := (C.int)(len(buf))
   a := make([]int8, nx)

   c_a := (*C.char)(unsafe.Pointer(&a[0]))
   c_buf := (*C.uchar)(unsafe.Pointer(&buf[0]))
   if int(C.fits_rdecomp_byte(c_buf, clen, c_a, C.int(nx), C.int(nblock))) != 0 {
      return nil
   }
   return a
}

func Fits_rdecomp_short(buf []uint8, nx int, nblock int) []int16 {
   clen := (C.int)(len(buf))
   a := make([]int16, nx)

   c_a := (*C.short)(unsafe.Pointer(&a[0]))
   c_buf := (*C.uchar)(unsafe.Pointer(&buf[0]))
   if int(C.fits_rdecomp_short(c_buf, clen, c_a, C.int(nx), C.int(nblock))) != 0 {
      return nil
   }
   return a
}
