package progress

import (
	"fmt"

	"github.com/prathyushnallamothu/dbconnection"

	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)



func GetProgress(cookieid string) float64 {
	var   Location, Typeofemp, Hourlyrate, Aboutme,  Skills, Collegename, Degree, CLocation, Startyear, Endyear, Certificationid, Certificationname, Certificationcompany, ImgURL string
	var totalprogress,aboutprogress,expprogress,eduprogress,skillprogress,certprogress,imgprogress,empprogress =0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0
	//var Companyname, Jobtitle, Jlocation, Startdate, Enddate, Technologies, Jobdescription, Projecttitle, Projecturl, Projectdescription float64
	db := dbconnection.Connect()
	result2, err1 := db.Query("select location,typeofemp,hourlyrate,aboutme from about where userid=?", cookieid)
	if err1 != nil {
		log.Fatal(err1)
	}
	for result2.Next() {
		err1 = result2.Scan( &Location, &Typeofemp, &Hourlyrate, &Aboutme)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
	
	if Location!=""{
		aboutprogress+=5
	}
	if Typeofemp!=""{
		aboutprogress+=5
	}
	if Hourlyrate!=""{
		aboutprogress+=5
	}
	if Aboutme!=""{
		aboutprogress+=5
	}
	result3, err2 := db.Query("select imagepath from images where userid=?", cookieid)
	if err2 != nil {
		log.Fatal(err2)
	}
	for result3.Next() {
		err2 = result3.Scan(&ImgURL)
		if err2 != nil {
			log.Fatal(err2)
		}
	}
	if ImgURL!=""{
		imgprogress+=10
	}
	var EmpCount int
	result, err := db.Query("select count(*) from employment where userid=?", cookieid)
	if err != nil {
		log.Fatal(err)
	}
	for result.Next() {
		err = result.Scan(&EmpCount)
		if err != nil {
			log.Fatal(err)
		}
	}
	if EmpCount>=1{
		empprogress+=10
	}
	
	ExpCount:=0
	result4,err4:=db.Query("select count(*) from experience where userid=?",cookieid)
	if err1!=nil{
		log.Fatal(err1)
	}	
	for result4.Next(){
		err4=result4.Scan(&ExpCount)
		if err4!=nil{
			log.Fatal(err4)
		}
	}
	if ExpCount>=1{
		expprogress+=10
	}
	SkillCount :=0
	result5,err5:=db.Query("select count(*),Skills from skills where userid=?",cookieid)
	if err5!=nil{
		log.Fatal(err5)
	}	
	for result5.Next(){
		err5=result5.Scan(&SkillCount,&Skills)
		if err1!=nil{
			log.Fatal(err5)
		}
	}
	fmt.Println(Skills)
	str:=strings.Split(Skills,",")
	
	if SkillCount<10 && SkillCount>0{
	for i:=len(str);i>0;i--{
		skillprogress+=2
	}
}
	if SkillCount>=10{
		skillprogress+=20
	}
	var Count int
	result8,err8:=db.Query("select count(*) from education where userid=?",cookieid)
	if err8!=nil{
		log.Fatal(err8)
	}
	for result8.Next(){
		err8=result8.Scan(&Count)
	}
	if Count<3{
	result6,err6:=db.Query("select collegename,degree,location,startyear,endyear from education where userid=?",cookieid)
	if err1!=nil{
		log.Fatal(err1)
	}	
	for result6.Next(){
		err6=result6.Scan(&Collegename,&Degree,&CLocation,&Startyear,&Endyear)
		if err6!=nil{
			log.Fatal(err6)
		}
		if Collegename!=""{
			eduprogress+=1.33
		}
		if Degree!=""{
			eduprogress+=1.33
		}
		if CLocation!=""{
			eduprogress+=1.33
		}
		if Startyear!=""{
			eduprogress+=1.33
		}
		if Endyear!=""{
			eduprogress+=1.33
		}
	}
	
}
	if Count>=3{
		eduprogress+=20
		fmt.Println(eduprogress)
	}
	result9,err9:=db.Query("select count(*) from certifications where userid=?",cookieid)
	if err9!=nil{
		log.Fatal(err9)
	}
	for result9.Next(){
		err9=result9.Scan(&Count)
		if err9!=nil{
			log.Fatal(err9)
		}
	}
	if Count<=2{
	result7,err7:=db.Query("select certificationid,certificationname,certificationcompany from certifications where userid=?",cookieid)
	if err7!=nil{
		log.Fatal(err7)
	}	
	for result7.Next(){
		err7=result7.Scan(&Certificationid,&Certificationname,&Certificationcompany)
		if err7!=nil{
			log.Fatal(err7)
		}
	}
	if Certificationid!=""{
		certprogress+=1.66
	}
	if Certificationname!=""{
		certprogress+=1.66
	}
	if Certificationcompany!=""{
		certprogress+=1.66
	}
}
	if Count>2{
		certprogress+=10
	}
	totalprogress=aboutprogress+expprogress+eduprogress+skillprogress+certprogress+imgprogress+empprogress
	return totalprogress
}
