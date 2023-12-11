import {FunctionComponent} from 'react';
import {EventDTO} from "../api/dtos/output.ts";
import {format} from 'date-fns';
import {ArrowRight} from "../icons/ArrowRight.tsx";
import {idEventShowing, showDetailsModalSubject} from "../utils/subjectsRx.ts";


type Props = EventDTO;

const EventCard: FunctionComponent<Props> = (event) => {
    const dateFormatted = new Date(event.Date)
    return (
        <div className="py-8 px-3  lg:w-1/3">
            <div className="py-8 px-3 h-full shadow-2xl  rounded-2xl">
                <div className="h-full flex items-start">
                    <div className="w-12 flex-shrink-0 flex flex-col text-center leading-none">
                    <span
                        className="text-gray-500 pb-2 mb-2 border-b-2 border-gray-200">{format(dateFormatted, 'MMM')}</span>
                        <span
                            className="font-medium text-lg text-gray-800 title-font leading-none">{format(dateFormatted, 'd')}</span>
                    </div>
                    <div className="flex-grow pl-6">
                        <h2 className="tracking-widest text-xs title-font font-medium text-indigo-500 mb-1">{event.Place.City}</h2>
                        <h1 className="title-font text-xl font-medium text-gray-900 mb-3">{event.Title}</h1>
                        <p className="leading-relaxed mb-5">{event.ShortDescription}</p>
                        <a className="inline-flex items-center">
                            <img alt="blog"
                                 src="https://static.vecteezy.com/system/resources/thumbnails/009/734/564/small/default-avatar-profile-icon-of-social-media-user-vector.jpg"
                                 className="w-8 h-8 rounded-full flex-shrink-0 object-cover object-center"/>
                            <span className="flex-grow flex flex-col pl-3">
                                    <span
                                        className="title-font font-medium text-gray-900">{event.Organizer.FirstName + " " + event.Organizer.LastName}</span>
                                  </span>
                            <a onClick={() => {
                                showDetailsModalSubject.setSubject(true)
                                idEventShowing.setSubject(event.Id)
                            }} className="text-indigo-500 pl-3 inline-flex items-center  cursor-pointer">Learn
                                More <ArrowRight/></a>

                        </a>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default EventCard;
