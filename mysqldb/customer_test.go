package mysqldb

import ("testing"
	"github.com/magiconair/properties/assert"
	"fmt"
	"github.com/mrtomyum/nopadol/mysqldb/n9model"
	"github.com/mrtomyum/nopadol/customer"
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

func Test_Search_by_id(t *testing.T){
	db,_ := connectDemo("npdl")
	cusRepo := NewCustomerRepository(db)
	r := customer.SearchByIdTemplate{Id:int64(53672)}

	cusRepo.SearchById(&r)
	rs , err := cusRepo.SearchById(r)

	if err != nil {
		t.Fatalf("error get id : %v",err.Error())
	}

	field, _ := rs.(*n9model.Customer)

	assert.Equal(t,field.Code,"41054")
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