package main

import (
	"git.changhong.io/qiniu-study/liwei/stringutils"
	"log"
)
// test stringutils 
func main() {

	log.Println(stringutils.VerifyEmail("contact@idrmfly.com"))
	log.Println(stringutils.VerifyEmail("contact.liw@idrmfly.com"))
	log.Println(stringutils.VerifyEmail("contact-liw@idrmfly.com"))
	log.Println(stringutils.VerifyEmail("contact-liwidrmfly.com"))
	log.Println(stringutils.VerifyEmail("contact-liw@idrmflycom"))

}
