package deque

import "testing"

func TestDeque(t *testing.T) {
	d := NewDeque()
	d.Push()
		t.Run(tt.name, func(t *testing.T) {
			d := &Deque{
				items: tt.fields.items,
				head:  tt.fields.head,
				tail:  tt.fields.tail,
				size:  tt.fields.size,
			}
			d.Push(tt.args.val)
		})
	}
}
