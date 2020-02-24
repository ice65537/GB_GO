package phonebook

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

//PhoneRecord -
type PhoneRecord struct {
	ContactName string
	Phones      []int
	Email       string
}

//PhoneBook -
type PhoneBook struct {
	recordset []PhoneRecord
	IdxName   map[string]int
}

func (x *PhoneBook) Len() int { return len(x.recordset) }
func (x *PhoneBook) Swap(i, j int) {
	x.IdxName[x.recordset[i].ContactName] = j
	x.IdxName[x.recordset[j].ContactName] = i
	//
	x.recordset[i], x.recordset[j] = x.recordset[j], x.recordset[i]
}
func (x *PhoneBook) Less(i, j int) bool {
	return x.recordset[i].ContactName < x.recordset[j].ContactName
}
func (x PhoneBook) String() string {
	var s string
	for _, value := range x.recordset {
		s += fmt.Sprintf("{%s (%s), [%d] номеров: ", value.ContactName, value.Email, len(value.Phones)) +
			fmt.Sprint(value.Phones) + "}\n"
	}
	return s
}

//LoadFromFile -
func (x *PhoneBook) LoadFromFile(filename string) error {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	//
	err = json.Unmarshal(buffer, &x.recordset)
	if err != nil {
		return err
	}
	x.resetIdx()
	//
	return nil
}

//SaveToFile -
func (x PhoneBook) SaveToFile(filename string) error {
	//
	buffer, err := json.Marshal(x.recordset)
	if err != nil {
		return err
	}
	//
	err = ioutil.WriteFile(filename, buffer, 0644)
	if err != nil {
		return err
	}
	return nil
}

//SetItem -
func (x *PhoneBook) SetItem(fio string, email string, pnumset ...int) {
	var rec PhoneRecord
	rec.ContactName = fio
	rec.Email = email
	rec.Phones = pnumset
	x.initIdx()
	if i, ok := x.IdxName[fio]; !ok {
		x.recordset = append(x.recordset, rec)
		x.IdxName[fio] = len(x.recordset) - 1
	} else {
		x.recordset[i] = rec
	}
}

//DelItem -
func (x *PhoneBook) DelItem(fio string) error {
	var newRecSet []PhoneRecord

	x.initIdx()
	if _, ok := x.IdxName[fio]; !ok {
		return errors.New(fmt.Sprint("Не найдена запись с индексом [", fio, "]"))
	}
	if x.IdxName[fio] > 0 {
		newRecSet = x.recordset[0:x.IdxName[fio]]
	}
	if x.IdxName[fio] < len(x.recordset)-1 {
		newRecSet = append(newRecSet, x.recordset[x.IdxName[fio]+1:]...)
	}
	x.recordset = newRecSet
	x.resetIdx()
	return nil
}

func (x *PhoneBook) initIdx() {
	if x.IdxName == nil {
		x.IdxName = make(map[string]int)
		for idx, value := range x.recordset {
			x.IdxName[value.ContactName] = idx
		}
	}
}

func (x *PhoneBook) resetIdx() {
	x.IdxName = nil
	x.initIdx()
}
