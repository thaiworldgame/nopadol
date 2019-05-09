package n9model

import (
	"testing"
	"github.com/mrtomyum/nopadol/mysqldb"
	//"github.com/magiconair/properties/assert"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func Test_customer_get(t *testing.T){
	testDB, err := mysqldb.ConnectDB("npdl")
	if err != nil {
		t.Fatal(err.Error())
	}

	e := "41054"
	c := Customer{}
	c.Get(testDB,e)
	fmt.Print(c)
	assert.Equal(t,c.Code,e)

}

func Test_customer_get_id_by_code(t *testing.T){
	testDB, err := mysqldb.ConnectDB("npdl")
	if err != nil {
		t.Fatal(err.Error())
	}
	e := "41054"
	c := Customer{}
	cid,_ := c.GetIdByCode(testDB,e)
	fmt.Print(cid)
	assert.Equal(t,cid,int64(53672))
}

func Test_customer_get_by_id(t *testing.T){
	testDB, err := mysqldb.ConnectDB("npdl")
	if err != nil {
		t.Fatal(err.Error())
	}
	e := int64(53672)
	c := Customer{}
	c.GetById(testDB,e)
	fmt.Println(c.Code)
	assert.Equal(t,c.Code,"41054")
}

func Test_customer_add_existingCustomer_must_fail(t *testing.T){
	testDB, err := mysqldb.ConnectDB("npdl")
	if err != nil {
		t.Fatal(err.Error())
	}
	e := int64(53672)
	c := Customer{}
	//get exists data from db
	c.GetById(testDB,e)
	err = c.Add(testDB)
	fmt.Println(err.Error())


	//must fail if add exists data
	assert.NotNil(t,err)
}



func Test_customer_update_not_exists_customer_must_fail(t *testing.T){
	testDB, err := mysqldb.ConnectDB("npdl")
	if err != nil {
		t.Fatal(err.Error())
	}
	e := int64(99999999)
	c := Customer{}
	//get exists data from db
	c.GetById(testDB,e)
	err = c.Update(testDB)
	fmt.Println(err.Error())


	//must fail if add exists data
	assert.NotNil(t,err)
}


func Test_customer_inactive(t *testing.T){
	//e := false
	testDB, err := mysqldb.ConnectDB("npdl")
	if err != nil {
		t.Fatal(err.Error())
	}
	r := int64(53672)
	c := Customer{}
	//get exists data from db
	_,err = c.Inactive(testDB,r)
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t,err)
}

func Test_customer_checkExists_by_code(t *testing.T){
	//e := false
	testDB, err := mysqldb.ConnectDB("npdl")
	if err != nil {
		t.Fatal(err.Error())
	}
	r := "41054"
	c := Customer{}
	assert.Equal(t,c.CheckExistByCode(testDB,r),true)
}

func Test_customer_checkExists_by_id(t *testing.T){
	//e := false
	testDB, err := mysqldb.ConnectDB("npdl")
	if err != nil {
		t.Fatal(err.Error())
	}
	r := int64(53672)
	c := Customer{}

	assert.Equal(t,c.CheckExistById(testDB,r),true)
}


func Test_customer_changecode(t *testing.T){
	testDB, err := mysqldb.ConnectDB("npdl")
	if err != nil {
		t.Fatal(err.Error())
	}
	r := "41054"
	c := Customer{}
	assert.Nil(t,err,c.ChangeCode(testDB,r))
}