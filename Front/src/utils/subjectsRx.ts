import {Subject} from "rxjs";

class SubscribeManagerRx<T> {
    private subject = new Subject<T>();

    public get getSubject() {
        return this.subject.asObservable();
    }

    public setSubject(value: T) {
        this.subject.next(value);
    }
}

export const showDetailsModalSubject = new SubscribeManagerRx<boolean>()
export const alertLoginModalSubject = new SubscribeManagerRx<boolean>()
export const idEventShowing = new SubscribeManagerRx<number>()
export const idEventEdit = new SubscribeManagerRx<number>()
export const pageToSkipSubject = new SubscribeManagerRx<number>()
