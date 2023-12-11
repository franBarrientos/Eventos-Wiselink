import { FunctionComponent, useEffect, useState } from "react"
import EventCard from "../components/EventCard.tsx"
import { CustomModal } from "../components/CustomModal.tsx"
import ModalDetailEvent from "../components/ModalDetailEvent.tsx"
import { GetMyEvents } from "../api/services/EventsService.ts"
import { toast } from "react-toastify"
import { EventDTO } from "../api/dtos/output.ts"
import { useNavigate } from "react-router-dom"
import {
  alertLoginModalSubject,
  idEventShowing,
  showDetailsModalSubject,
} from "../utils/subjectsRx.ts"
import { FieldValues, SubmitHandler, useForm } from "react-hook-form"
import CardSkeleton from "../components/CardSkeleton.tsx";

interface OwnProps {}

type Props = OwnProps

const MyEvents: FunctionComponent<Props> = () => {
  const [events, setEvents] = useState<EventDTO[]>()
  const [detailsModalIsOpen, setDetailsModalIsOpen] = useState(false)
  const [idDetailsModalOpen, setIdDetailsModalOpen] = useState<number>(0)
  const [alertLoginModal, setAlertLoginModal] = useState(false)
  const navigate = useNavigate()
  const { register, handleSubmit } = useForm()
  const [isLoading, setIsLoading] = useState(true)

  //hoks
  useEffect(() => {
    GetMyEvents("")
      .then((r) => {
        setEvents(r.eventsSubscribed)
        setTimeout(() => {
            setIsLoading(false)
        }, 1000)
      })
      .catch(() => {
        toast("Something went wrong")
      })
  }, [])

  //RxSubjects
  useEffect(() => {
    const subscribe = showDetailsModalSubject.getSubject.subscribe((value) => {
      setDetailsModalIsOpen(value)
    })

    const subscribe2 = idEventShowing.getSubject.subscribe((value) => {
      setIdDetailsModalOpen(value)
    })

    const subscribe3 = alertLoginModalSubject.getSubject.subscribe((value) => {
      setAlertLoginModal(value)
    })

    return () => {
      subscribe.unsubscribe()
      subscribe2.unsubscribe()
      subscribe3.unsubscribe()
    }
  }, [detailsModalIsOpen, idDetailsModalOpen, alertLoginModal])

  const handleFilter = (data: { state: string }) => {
    setIsLoading(true)
    GetMyEvents(data.state)
      .then((r) => {
        setEvents(r.eventsSubscribed)
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
            My Events
          </h1>
          <form
            onSubmit={handleSubmit(handleFilter as SubmitHandler<FieldValues>)}
            className="flex justify-center my-8 items-center space-x-4"
          >


            <label className="mr-2">
              <input
                type="radio"
                value="active"
                className="mr-1"
                {...register("state")}
              />
              <span>Active</span>
            </label>

            <label>
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
            { !isLoading && events?.map((event) => <EventCard {...event} />)}
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

export default MyEvents
