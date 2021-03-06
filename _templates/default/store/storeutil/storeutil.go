// Package storeutil contains a ListenerRegistry type. Ripped from vecty examples with a few modifications.
package storeutil

// ListenerRegistry is a listener registry.
// The zero value is unfit for use; use NewListenerRegistry to create an instance.
type ListenerRegistry struct {
	listeners map[interface{}]func(action interface{})
}

// NewListenerRegistry creates a listener registry.
func NewListenerRegistry() *ListenerRegistry {
	return &ListenerRegistry{
		listeners: make(map[interface{}]func(interface{})),
	}
}

// Add adds listener with key to the registry.
// key may be nil, then an arbitrary unused key is assigned.
// It panics if a listener with same key is already present.
func (r *ListenerRegistry) Add(key interface{}, listener func(action interface{})) {
	if key == nil {
		key = new(int)
	}
	if _, ok := r.listeners[key]; ok {
		panic("duplicate listener key")
	}
	r.listeners[key] = listener
}

// Remove removes a listener with key from the registry.
func (r *ListenerRegistry) Remove(key interface{}) {
	delete(r.listeners, key)
}

// Fire invokes all listeners in the registry.
func (r *ListenerRegistry) Fire(action interface{}) {
	for _, l := range r.listeners {
		l(action)
	}
}
