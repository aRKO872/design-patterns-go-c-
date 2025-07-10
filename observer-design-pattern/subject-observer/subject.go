package subject_observer

type Order struct {
	ID	 string
	Observers map[string]Observer
}

func (o *Order) AddObserver(observer Observer) {
	if o.Observers == nil {
		o.Observers = make(map[string]Observer)
	}
	o.Observers[observer.GetName()] = observer
}

func (o *Order) RemoveObserver(observer Observer) {
	if o.Observers != nil {
		delete(o.Observers, observer.GetName())
	}
}

func (o *Order) NotifyObservers() {
	if o.Observers != nil {
		for _, observer := range o.Observers {
			observer.SendNotification(*o)
		}
	}
}