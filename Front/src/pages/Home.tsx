import {useEffect, useState} from "react"
import {getAllEventsPublished, GetEventsFiltered,} from "../api/services/EventsService.ts"
import {toast} from "react-toastify"
import {EventDTO} from "../api/dtos/output.ts"
import EventCard from "../components/EventCard.tsx"
import {CustomModal} from "../components/CustomModal.tsx"
import {
    alertLoginModalSubject,
    idEventShowing,
    pageToSkipSubject,
    showDetailsModalSubject,
} from "../utils/subjectsRx.ts"
import ModalDetailEvent from "../components/ModalDetailEvent.tsx"
import {useNavigate} from "react-router-dom"
import {FieldValues, SubmitHandler, useForm} from "react-hook-form"
import CardSkeleton from "../components/CardSkeleton.tsx";
import 'react-loading-skeleton/dist/skeleton.css'
import Pagination from "../components/Pagination.tsx";

export default function Home() {
    const [events, setEvents] = useState<EventDTO[]>([])
    const [detailsModalIsOpen, setDetailsModalIsOpen] = useState(false)
    const [idDetailsModalOpen, setIdDetailsModalOpen] = useState<number>(0)
    const [alertLoginModal, setAlertLoginModal] = useState(false)
    const [isLoading, setIsLoading] = useState(true)
    const {register, handleSubmit, getValues} = useForm()
    const [currentPage, setCurrentPage] = useState(1)

    const [isFilter, setIsFilter] = useState(false)


    //hoks
    const navigate = useNavigate()

    const getAllEvents = (id: number = 1) => {
        getAllEventsPublished(id)
            .then((r) => {
                setEvents(r)
                setTimeout(() => {
                    setIsLoading(false)
                }, 1000)
            })
            .catch(() => {
                toast("Something went wrong")
            })
    }

    useEffect(() => {
        getAllEvents()
    }, [])

    //RxSubjects
    useEffect(() => {
        const subscribeShowDetail = showDetailsModalSubject.getSubject.subscribe((value) => {
            setDetailsModalIsOpen(value)
        })

        const subscribeIdEvent = idEventShowing.getSubject.subscribe((value) => {
            setIdDetailsModalOpen(value)
        })

        const subscribeAlertModal = alertLoginModalSubject.getSubject.subscribe((value) => {
            setAlertLoginModal(value)
        })

        const listenToCurrentPage = pageToSkipSubject.getSubject.subscribe((value) => {
            setCurrentPage(value)
            if (isFilter) {
                console.log({ isFilter })
                const data = getValues();
                const dateFormatted =
                    data.date != "" ? new Date(data.date).toISOString() : ""
                GetEventsFiltered(data.title, dateFormatted, data.state, value)
                    .then((r) => {
                        setEvents(r)
                        setTimeout(() => {
                            setIsLoading(false)
                        }, 1000)
                    })
                    .catch(() => {
                        toast("Something went wrong")
                    })

                } else {
                getAllEvents(value)
                }
            }
        )


        return () => {
            subscribeShowDetail.unsubscribe()
            subscribeIdEvent.unsubscribe()
            subscribeAlertModal.unsubscribe()
            listenToCurrentPage.unsubscribe()
        }
    }, [detailsModalIsOpen, idDetailsModalOpen, alertLoginModal, isFilter])


console.log({isFilterDespues : isFilter})


    const handleFilter = (data: {
        title: string
        date: string
        state: string
    }) => {
        setIsLoading(true)
        setCurrentPage(1)
        setIsFilter(true)
        const dateFormatted =
            data.date != "" ? new Date(data.date).toISOString() : ""
        GetEventsFiltered(data.title, dateFormatted, data.state, 1)
            .then((r) => {
                setEvents(r)
                setTimeout(() => {
                    setIsLoading(false)
                }, 1000)
            })
            .catch(() => {
                toast("Something went wrong")
            })
    }


    return (
        <>
            <section className="text-gray-600 body-font">
                <div className="container px-5 py-8 mx-auto">
                    <h1 className="text-center mb-5 sm:text-3xl text-2xl font-medium title-font text-gray-900">
                        Published Events
                    </h1>
                    <form
                        onSubmit={handleSubmit(handleFilter as SubmitHandler<FieldValues>)}
                        className="flex flex-col gap-5 md:gap-1 md:flex-row justify-center my-8 items-center space-x-4"
                    >
                        {/* Filtro por nombre */}
                        <input
                            type="text"
                            placeholder="Title"
                            className="ml-3 rounded-xl border w-1/2 md:w-auto px-2 py-1 rounded"
                            {...register("title")}
                        />

                        {/* Filtro por fecha */}
                        <input
                            type="date"
                            className="rounded-xl border w-1/2 md:w-auto px-2 py-1 rounded"
                            {...register("date")}
                        />

                        <div className="flex gap-1">
                            <label className="mr-2">
                                <input
                                    type="radio"
                                    value="active"
                                    className="mr-1"
                                    {...register("state")}
                                />
                                <span>Active</span>
                            </label>

                            <label className="mr-2">
                                <input
                                    type="radio"
                                    value="completed"
                                    className="mr-1"
                                    {...register("state")}
                                />
                                <span>Finished</span>
                            </label>

                            <label className="mr-2">
                                <input
                                    type="radio"
                                    value=""
                                    className="mr-1"
                                    {...register("state")}
                                />
                                <span>Both</span>
                            </label>
                        </div>

                        {/* Botón para aplicar el filtro */}
                        <button
                            type={"submit"}
                            className="bg-indigo-500 text-white px-3 py-1 rounded"
                        >
                            Filter
                        </button>
                    </form>
                    <div className="flex flex-wrap -mx-4  -my-8">
                        {isLoading && (
                            <>
                            <CardSkeleton />
                            <CardSkeleton />
                            <CardSkeleton />
                            <CardSkeleton />
                            <CardSkeleton />
                            <CardSkeleton />
                            <CardSkeleton />
                            <CardSkeleton />
                            <CardSkeleton />
                            <CardSkeleton />
                            <CardSkeleton />
                            <CardSkeleton />

                            </>
                            )}
                        { !isLoading && events.map((event) => <EventCard {...event} />)}

                    </div>
                    <div className="mx-auto w-full flex justify-center mt-10">
                        <Pagination currentPage={currentPage} />
                    </div>
                </div>
            </section>

            <CustomModal
                isOpen={detailsModalIsOpen}
                label={"Details"}
                onRequestClose={() => setDetailsModalIsOpen(false)}
            >
                {alertLoginModal ? (
                    <div className="fixed inset-0 flex items-center justify-center z-50">
                        <div className="bg-white p-6 rounded-md shadow-md text-center">
                            <h1 className="text-2xl font-bold mb-4">Sign In Required</h1>
                            <p className="text-gray-600 mb-4">
                                To register for this event, please sign in to your account.
                            </p>
                            <div className="flex justify-center">
                                <button
                                    className="bg-indigo-500 hover:bg-blue-700 text-white font-bold py-2 px-4 mr-4 rounded focus:outline-none focus:shadow-outline-blue"
                                    onClick={() => {
                                        navigate("/login")
                                    }}
                                >
                                    Sign In
                                </button>
                                <button
                                    className="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline-gray"
                                    onClick={() => {
                                        setAlertLoginModal(false)
                                    }}
                                >
                                    No
                                </button>
                            </div>
                        </div>
                    </div>
                ) : events && events.length > 0 ? (
                    events.find((e) => e.Id === idDetailsModalOpen) ? (
                        <ModalDetailEvent
                            {...events.find((e) => e.Id === idDetailsModalOpen)!}
                        />
                    ) : (
                        <div>Event with id {idDetailsModalOpen} not found.</div>
                    )
                ) : (
                    <div>No events found.</div>
                )}
            </CustomModal>
        </>
    )
}
