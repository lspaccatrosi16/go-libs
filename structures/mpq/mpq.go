package mpq

type Item[T any] struct {
	Priority int
	Val      T
}

type Queue[T any] struct {
	Items []Item[T]
}

func (q *Queue[T]) Add(val T, priority int) {
	item := Item[T]{
		Priority: priority,
		Val:      val,
	}
	if len(q.Items) == 0 {
		q.Items = append(q.Items, item)
		return
	}

	found := false

	for i, v := range q.Items {
		if item.Priority > v.Priority {
			if i == 0 {
				q.Items = append([]Item[T]{item}, q.Items...)
			} else {
				q.Items = append(q.Items[:i], q.Items[i+1:]...)
				q.Items[i] = item
			}
			found = true
			break
		}
	}

	if !found {
		q.Items = append(q.Items, item)
	}
}

func (q *Queue[T]) Pop() T {
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item.Val
}

func (q *Queue[T]) Len() int {
	return len(q.Items)
}
