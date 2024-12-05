type Type = string
type Payload = {}
type Listener = (_: Payload) => any

class EventDetail {
  type: Type;
  payload: Payload;
}

const events: { [key in Type]: Listener[] } = {};
const ErrAlreadyOccupied = (type: Type) => new Error(`event "${type}" is already occupied. use other event name`);
const ErrDoNotExists = (type: Type) => new Error(`event "${type}" does not exist. Use EventListener occupy to create this event`);

class EventListener {
  static Dispach(event: EventDetail) {
    if (!events[event.type])
      throw ErrDoNotExists(event.type);

    setTimeout(() =>
      events[event.type].map(fn => fn(event.payload))
    );
  }

  static AddListener(type: Type, func: Listener, first: boolean = false) {
    if (!events[type])
      throw ErrDoNotExists(type);

    first ?
      events[type].unshift(func) :
      events[type].push(func);
  }

  static RemoveListener(type: Type, func: Listener) {
    if (!events[type])
      throw ErrDoNotExists(type);
    events[type] = events[type].filter(e => e != func)
  }

  static Occupy(type: Type) {
    if (events[type])
      throw ErrAlreadyOccupied(type);
    events[type] ??= [];
  }
}





class CreateEventOutput<T extends Payload> {
  Dispatch: (payload: T) => void
  AddListener: (func: (_: T) => any, first?: boolean) => void
  RemoveListener: (func: (_: T) => any) => void
}

const CreateEvent = <T extends Payload>(type: Type): CreateEventOutput<T> => {
  EventListener.Occupy(type)
  return {
    Dispatch(payload) {
      EventListener.Dispach({
        type: type,
        payload: payload
      })
    },
    AddListener(func, first = false) { // @ts-ignore
      EventListener.AddListener(type, func, first)
    },
    RemoveListener(func) { // @ts-ignore
      EventListener.RemoveListener(type, func)
    },
  }
}

export {
  CreateEvent
}