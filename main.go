package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	var nama string
	fmt.Print("Masukkan nama anda : ")
	fmt.Scanln(&nama)
	fmt.Printf("Baik nama anda adalah %s\n", nama)
}

//message
/*
- git init
- git status
- rm -rf
- git add . / git add name_file
- git commit -m "message"
- git commit -am "message"
- example message
		[FEATURE] untuk penambahan fitur
		[IMPROVE] jika ada perubahan
		[FIX] untuk project yang fix
		[MERGE]
- git log untuk melihat riwayat git
- git log --oneline
- git branch nama_branch
- git checkout nama_branch
- git merge nama_branch_yang_mau_dimerge
- git branch nama_branch_baru branch_tujuan walaupun dalam branch sekarang
- git clone
- git cherry-pic commitID cara mendapatkan commitID dengan sara git log
- git cherry-pick --continue melanjutkan
- git cherry-pick --abort membatalkan
- git cherry-pick --skip untuk melompati commitID ke commitID selanjutnya

*/
