import { useEffect, useState } from "react"
import {
  CreateEvent,
  EditService,
  GetEventsAdmin,
  GetEventsFiltered,
} from "../api/services/EventsService.ts"
import { toast } from "react-toastify"
import { EventDTO, UserDTO } from "../api/dtos/output.ts"
import EventCard from "../components/EventCard.tsx"
import { CustomModal } from "../components/CustomModal.tsx"
import {
  alertLoginModalSubject,
  idEventEdit,
  idEventShowing,
  showDetailsModalSubject,
} from "../utils/subjectsRx.ts"
import ModalDetailEventAdmin from "../components/ModalDetailEventAdmin.tsx"
import { useNavigate } from "react-router-dom"
import { FieldValues, SubmitHandler, useForm } from "react-hook-form"
import {EventAddDTO} from "../api/dtos/input.ts";

export default function Admin() {
  const navigate = useNavigate()

  const [events, setEvents] = useState<EventDTO[]>()
  const [detailsModalIsOpen, setDetailsModalIsOpen] = useState(false)
  const [idDetailsModalOpen, setIdDetailsModalOpen] = useState<number>(0)
  const [editModal, setEditModal] = useState(false)
  const [createModal, setCreateModal] = useState(false)

  const { register, handleSubmit } = useForm()
  const { register: registerEdit, handleSubmit: handleSubmitEdit, reset } = useForm()
  const { register: registerCreate, handleSubmit: handleSubmitCreate } =
    useForm()
  const [eventToEdit, setEventToEdit] = useState<EventDTO>()

  useEffect(() => {
    if (!localStorage.getItem("token")) {
      toast.error("You don't have permission to access this page")
      navigate("/")
      return
    }

    if (
      (JSON.parse(localStorage.getItem("user") ?? "") as UserDTO).Role !=
      "ADMIN"
    ) {
      toast.error("You don't have permission to access this page")
      navigate("/")
      return
    }

    console.log(localStorage.getItem("token"))

    GetEventsAdmin()
      .then((r) => {
        setEvents(r)
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
      setEditModal(value)
    })

    const subscribe4 = idEventEdit.getSubject.subscribe((value) => {
      reset()
      setEventToEdit(events?.find((e) => e.Id == value))
    })

    return () => {
      subscribe.unsubscribe()
      subscribe2.unsubscribe()
      subscribe3.unsubscribe()
      subscribe4.unsubscribe()
    }
  }, [detailsModalIsOpen, idDetailsModalOpen, editModal, eventToEdit])

  const handleFilter = (data: {
    title: string
    date: string
    state: string
  }) => {
    const dateFormatted =
      data.date != "" ? new Date(data.date).toISOString() : ""
    GetEventsFiltered(data.title, dateFormatted, data.state)
      .then((r) => {
        console.log(r)
        setEvents(r)
      })
      .catch(() => {
        toast("Something went wrong")
      })
  }

  const editEvent = (data: EventAddDTO) => {
    console.log(data)
    EditService(eventToEdit!.Id!, data)
      .then((r) => {
        toast.success("Event edited successfully")
        setEvents(
          events?.map((e) => {
            if (e.Id == r.Id) {
              return r
            }
            return e
          }),
        )
        setDetailsModalIsOpen(false)
        setEditModal(false)
      })
      .catch((e) => {
        toast.error(e.toString())
      })
  }

  const createEvent = (data: EventAddDTO) => {
    CreateEvent({
      ...data,
      Place: { ...data.Place, AddressNumber: Number(data.Place.AddressNumber) },
      Date: new Date(data.Date).toISOString(),
    })
      .then((r) => {
        toast.success("Event added successfully")
        setEvents([...events!, r])
        setDetailsModalIsOpen(false)
        setEditModal(false)
        setCreateModal(false)
      })
      .catch((e) => {
        toast.error(e.toString())
      })
  }

  return (
    <>
      <section className="text-gray-600 body-font">
        <div className="container px-5 py-8 mx-auto">
          <h1 className="text-center mb-5 sm:text-3xl text-2xl font-medium title-font text-gray-900">
            Published Events
          </h1>
          <div className="flex flex-col md:flex-row w-full items-center">
            <button
              onClick={() => {
                setDetailsModalIsOpen(true)
                setCreateModal(true)
              }}
              type={"submit"}
              className="bg-indigo-500 text-white px-3 py-1 rounded"
            >
              Create Event
            </button>
            <form
              onSubmit={handleSubmit(
                handleFilter as SubmitHandler<FieldValues>,
              )}
              className="w-5/6 flex flex-col gap-5 md:gap-1 md:flex-row justify-center my-8 items-center space-x-4"
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
          </div>

          <div className="flex flex-wrap -mx-4  -my-8">
            {events?.map((event) => <EventCard {...event} />)}
          </div>
        </div>
      </section>

      <CustomModal
        isOpen={detailsModalIsOpen}
        label={"Details"}
        onRequestClose={() => {
          setDetailsModalIsOpen(false)
          setEditModal(false)
          setCreateModal(false)
          setEventToEdit(undefined)
        }}
      >
        {createModal ? (
          <div className="fixed inset-0 flex items-center justify-center z-50">
            <div className="bg-white p-6 rounded-md shadow-md text-center">
              <h1 className="text-2xl font-bold mb-4">Create Event</h1>
              <form
                className="md:grid md:grid-cols-2 gap-5"
                onSubmit={handleSubmitCreate(createEvent as SubmitHandler<FieldValues>)}
              >
                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Title
                  </label>
                  <input
                    {...registerCreate("Title", {
                      required: true,
                    })}
                    type="text"
                    className="border rounded-lg px-3 py-2  text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Short Description
                  </label>
                  <input
                    {...registerCreate("ShortDescription", {
                      required: true,
                    })}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className=" w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Long Description
                  </label>
                  <input
                    {...registerCreate("LongDescription", {
                      required: true,
                    })}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Date
                  </label>
                  <input
                    {...registerCreate("Date", {
                      required: true,
                    })}
                    type="datetime-local"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className=" w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Organizer FirstName
                  </label>
                  <input
                    {...registerCreate("Organizer.FirstName", {
                      required: true,
                    })}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4  text-sm text-gray-600 pb-1 block">
                    Organizer LastName
                  </label>
                  <input
                    {...registerCreate("Organizer.LastName", {
                      required: true,
                    })}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Address
                  </label>
                  <input
                    {...registerCreate("Place.Address", {
                      required: true,
                    })}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>
                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Address Number
                  </label>
                  <input
                    {...registerCreate("Place.AddressNumber", {
                      required: true,
                    })}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    City
                  </label>
                  <input
                    {...registerCreate("Place.City", {
                      required: true,
                    })}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Country
                  </label>
                  <input
                    {...registerCreate("Place.Country.", {
                      required: true,
                    })}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div className="flex items-center justify-center">
                  <input
                    id="checkbox"
                    type="checkbox"
                    defaultChecked={true}
                    {...registerCreate("State")}
                    className="form-checkbox h-5 w-5 text-indigo-500 focus:ring-indigo-500 border-gray-300 rounded"
                  />
                  <label htmlFor="checkbox" className="ml-2 text-gray-700">
                    Active
                  </label>
                </div>

                <button
                  type="submit"
                  className="transition duration-200 bg-indigo-500 hover:bg-indigo-600 focus:bg-indigo-700
                   focus:shadow-sm focus:ring-4 focus:ring-blue-500 focus:ring-opacity-50
                    text-white w-full h-full  py-2.5 mt-1 rounded-lg text-sm shadow-sm hover:shadow-md font-semibold text-center inline-block"
                >
                  <span className="inline-block mr-2">Save</span>
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                    className="w-4 h-4 inline-block"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M17 8l4 4m0 0l-4 4m4-4H3"
                    />
                  </svg>
                </button>
              </form>
            </div>
          </div>
        ) : editModal ? (
          <div className="fixed inset-0 flex items-center justify-center z-50">
            <div className="bg-white p-6 rounded-md shadow-md text-center">
              <h1 className="text-2xl font-bold mb-4">Edit Event</h1>
              <form
                className="md:grid md:grid-cols-2 gap-5"
                onSubmit={handleSubmitEdit(editEvent as SubmitHandler<FieldValues>)}
              >
                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Title
                  </label>
                  <input
                    {...registerEdit("Title", {
                      required: true,
                    })}
                    defaultValue={eventToEdit?.Title}
                    type="text"
                    className="border rounded-lg px-3 py-2  text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Short Description
                  </label>
                  <input
                    {...registerEdit("ShortDescription", {
                      required: true,
                    })}
                    defaultValue={eventToEdit?.ShortDescription}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className=" w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Long Description
                  </label>
                  <input
                    {...registerEdit("LongDescription", {
                      required: true,
                    })}
                    defaultValue={eventToEdit?.LongDescription}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Date
                  </label>
                  <input
                    {...registerEdit("Date", {
                      required: true,
                    })}
                    defaultValue={new Date(eventToEdit!.Date!)
                      .toISOString()
                      .slice(0, 16)}
                    type="datetime-local"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className=" w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Organizer FirstName
                  </label>
                  <input
                    {...registerEdit("Organizer.FirstName", {
                      required: true,
                    })}
                    defaultValue={eventToEdit?.Organizer.FirstName}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4  text-sm text-gray-600 pb-1 block">
                    Organizer LastName
                  </label>
                  <input
                    {...registerEdit("Organizer.LastName", {
                      required: true,
                    })}
                    defaultValue={eventToEdit?.Organizer.LastName}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Address
                  </label>
                  <input
                    {...registerEdit("Place.Address", {
                      required: true,
                    })}
                    defaultValue={eventToEdit?.Place.Address}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>
                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Address Number
                  </label>
                  <input
                    {...registerEdit("Place.AddressNumber", {
                      required: true,
                    })}
                    defaultValue={eventToEdit?.Place.AddressNumber}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    City
                  </label>
                  <input
                    {...registerEdit("Place.City", {
                      required: true,
                    })}
                    defaultValue={eventToEdit?.Place.City}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div
                  className={"inline-flex justify-center items-center gap-2"}
                >
                  <label className="w-1/4 font-semibold text-sm text-gray-600 pb-1 block">
                    Country
                  </label>
                  <input
                    {...registerEdit("Place.Country.", {
                      required: true,
                    })}
                    defaultValue={eventToEdit?.Place.Country}
                    type="text"
                    className="border rounded-lg px-3 py-2 mt-1 mb-5 text-sm w-full"
                  />
                </div>

                <div className="flex items-center justify-center">
                  <input
                    id="checkbox"
                    type="checkbox"
                    {...registerEdit("State")}
                    defaultChecked={eventToEdit?.State}
                    className="form-checkbox h-5 w-5 text-indigo-500 focus:ring-indigo-500 border-gray-300 rounded"
                  />
                  <label htmlFor="checkbox" className="ml-2 text-gray-700">
                    Active
                  </label>
                </div>
                <button
                  type="submit"
                  className="transition duration-200 bg-indigo-500 hover:bg-indigo-600 focus:bg-indigo-700
                   focus:shadow-sm focus:ring-4 focus:ring-blue-500 focus:ring-opacity-50
                    text-white w-full col-span-2 h-full  py-2.5 mt-1 rounded-lg text-sm shadow-sm hover:shadow-md font-semibold text-center inline-block"
                >
                  <span className="inline-block mr-2">Save</span>
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                    className="w-4 h-4 inline-block"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      stroke-width="2"
                      d="M17 8l4 4m0 0l-4 4m4-4H3"
                    />
                  </svg>
                </button>
              </form>
            </div>
          </div>
        ) : events && events.length > 0 ? (
          events.find((e) => e.Id === idDetailsModalOpen) ? (
            <ModalDetailEventAdmin
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
