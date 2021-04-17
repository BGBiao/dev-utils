package list

import (
	"container/list"
)

type listStruct struct {
	l *list.List
}

func NewList() *listStruct {
	return &listStruct{l: list.New()}
}

func (l *listStruct) Len() int {
	return l.l.Len()
}

// init a new container/list object with int32
// mabey , you can use generate code.
func (l *listStruct) GenerateListInt32(data []int32) {
	firstValue := l.l.Front()
	if firstValue == nil {
		for i := 0; i < len(data); i++ {

			if i == 0 {
				l.l.PushFront(data[i])
			} else {

				// *list.Element
				lastV := l.l.Back()
				// fmt.Println(data[i], lastV)
				l.l.InsertAfter(data[i], lastV)
			}

		}

	} else {
		lastV := l.l.Back()
		for _, v := range data {
			l.l.InsertAfter(v, lastV)
		}

	}

}

func (l *listStruct) Pop() interface{} {
	lastV := l.l.Back()

	data := lastV.Value
	l.l.Remove(lastV)
	if lastV != nil {
		return data
	}

	return nil

}

func (l *listStruct) GetListValue() []interface{} {
	var datas []interface{}
	if l.l.Len() == 0 {
		return datas
	}

	// notice : l.l.Len() is 8 ,but the l.Len() is 4
	listLength := l.l.Len()
	for i := 0; i < listLength; i++ {
		lastV := l.l.Back()
		l.l.Remove(lastV)
		datas = append(datas, lastV.Value)

	}

	return datas
}
