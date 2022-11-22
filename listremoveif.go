package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	da := &List{}
	for i := l.Head; i != nil; i = i.Next {
		if i.Data != data_ref {
			ListPushBack(da, i.Data)
		}
	}
	l.Head = da.Head
}
