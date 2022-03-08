package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/timwhitez/Doge-EGGCall/pkg/eggcall"
	"github.com/timwhitez/Doge-EGGCall/pkg/eggreplace"
	"reflect"
	"unsafe"
)

func main(){
	//NtDelayExecution HellsGate
	sleep1,e := eggcall.MemHgate("84804f99e2c7ab8aee611d256a085cf4879c4be8",str2sha1)
	if e != nil {
		panic(e)
	}

	fmt.Printf("%s: %x\n","NtDelayExecution Sysid",sleep1)
	times := -(5000 * 10000)

	//hellsgate syscall

	eggreplace.FindAndReplace(
		[]byte{0x65,0x67,0x67,0x63,0x61,0x6c,0x6c},
		[]byte{0x90,0x90,0x0f,0x05,0x90,0x90,0x90},
		reflect.ValueOf(eggcall.HgSyscall).Pointer())

	eggcall.HgSyscall(sleep1,0,uintptr(unsafe.Pointer(&times)))
}




func str2sha1(s string) string{
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}


func Sha256Hex(s string)string{
	return hex.EncodeToString(Sha256([]byte(s)))
}

func Sha256(data []byte)[]byte{
	digest:=sha256.New()
	digest.Write(data)
	return digest.Sum(nil)
}
