package mysqldb

import ("testing"
	"github.com/magiconair/properties/assert"
	"fmt"
)


func Test_check_exits_by_code(t *testing.T){
	db,_ := connectDemo("npdl")
	c := CustomerModel{}
	c.Search(db,"41054")
	//_,err  := c.save(db)

	assert.Equal(t,c.Code,"41054")
}


func Test_get_id_by_code(t *testing.T){
	db,_ := connectDemo("npdl")
	c := CustomerModel{}
	id , err := c.getIdByCode(db,"41054")
	if err != nil {
		t.Fatalf("error get id : %v",err.Error())
	}
	var e int64
	e = 53672
	assert.Equal(t,id,e)
}


func Test_save_existing_customer_must_not_double_record(t *testing.T){
	db,_ := connectDemo("npdl")
	c := CustomerModel{}

	var count int
	e := "41054"
	c.Search(db,e)
	_,err := c.save(db)
	if err != nil {
		t.Fatalf("error query %v \n",err.Error())
	}

	lccommand := "select count(id) as countx from Customer where code='"+e+"'"
	fmt.Println(lccommand)

	err = db.Get(&count,lccommand)
	if err != nil {
		t.Fatalf("error query %v \n",err.Error())
	}
	assert.Equal(t,count,1)

}