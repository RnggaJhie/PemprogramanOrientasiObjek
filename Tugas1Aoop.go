package main

import "fmt"

type lagu struct{
    artis string;
    judul string;
    tanggalrilis int;
}

func main(){
 var kumpulan []lagu

    kumpulan = []lagu{
     lagu{
         artis: "Blackpink",
         judul : "lovesick girls",
         tanggalrilis: 2020,
     },
     lagu{
         artis: "blackpink",
         judul : "how you like that",
         tanggalrilis: 2020,
     },
     lagu{
         artis: "EXO",
         judul : "Lucky One",
         tanggalrilis: 2016,
     },
     lagu{
         artis: "EXO",
         judul : "Tempo",
         tanggalrilis: 2018,
     },
     lagu{
         artis: "BTS",
         judul: "the eternal",
         tanggalrilis: 2019,
     },
     lagu{
         artis: "BTS",
         judul : "Dionysius",
         tanggalrilis: 2019,
     },
     lagu{
         artis: "Melly",
         judul : "Mungkin",
         tanggalrilis: 2019,
     },
     lagu{
         artis: "Justin T",
         judul : "Mirrors",
         tanggalrilis: 2016,
     }, 
     lagu{
         artis: "EXO",
         judul : "For life",
         tanggalrilis: 2016,
     }, 
 }

 fmt.Println("-------Data Lagu--------")
 fmt.Println(kumpulan[0])
 fmt.Println(kumpulan[1])
 fmt.Println(kumpulan[2])
 fmt.Println(kumpulan[3])
 fmt.Println(kumpulan[4])
 fmt.Println(kumpulan[5])
 fmt.Println(kumpulan[6])
 fmt.Println(kumpulan[7])
 fmt.Println(kumpulan[8])
 fmt.Println(kumpulan[9])
 
}
