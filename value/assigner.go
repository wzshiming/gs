package value

type Assigner interface {
	Set(k Value, v Value)
	SetLocal(k Value, v Value)
	Get(k Value) (Value, bool)
	Child() Assigner
}
