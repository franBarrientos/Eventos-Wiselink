import { FunctionComponent } from "react"
import { EventDTO } from "../api/dtos/output.ts"
import {
  alertLoginModalSubject,
  idEventEdit,
  showDetailsModalSubject,
} from "../utils/subjectsRx.ts"
import { format } from "date-fns"

type Props = EventDTO

const ModalDetailEventAdmin: FunctionComponent<Props> = (event) => {
  return (
    <section className="text-gray-600 body-font">
      <div className="px-5 py-16 mx-auto">
        <div className="text-center mb-20">
          <h1 className="sm:text-3xl text-2xl font-medium text-center title-font text-gray-900 mb-4">
            {event.Title}
          </h1>
          <p className="text-base leading-relaxed xl:w-2/4 lg:w-3/4 mx-auto">
            {event.ShortDescription}
          </p>
        </div>
        <div className="text-center mb-20">
          <h1 className="sm:text-3xl text-2xl font-medium text-center title-font text-gray-900 mb-4">
            Description
          </h1>
          <p className="text-base leading-relaxed xl:w-2/4 lg:w-3/4 mx-auto">
            {event.LongDescription}
          </p>
        </div>
        <div className="flex flex-col gap-1 md:gap-0 md:flex-row  flex-nowrap justify-center">
          <div className="flex justify-start pb-2 flex-1 flex-col items-center  border-b   border-gray-200  ">
            <div className="sm:w-32 sm:order-none order-first sm:h-32 h-20 w-20  inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 flex-shrink-0">
              <svg
                className="sm:w-16 sm:h-16 w-10 h-10"
                aria-hidden="true"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 20 20"
              >
                <path
                  fill="currentColor"
                  d="M6 1a1 1 0 0 0-2 0h2ZM4 4a1 1 0 0 0 2 0H4Zm7-3a1 1 0 1 0-2 0h2ZM9 4a1 1 0 1 0 2 0H9Zm7-3a1 1 0 1 0-2 0h2Zm-2 3a1 1 0 1 0 2 0h-2ZM1 6a1 1 0 0 0 0 2V6Zm18 2a1 1 0 1 0 0-2v2ZM5 11v-1H4v1h1Zm0 .01H4v1h1v-1Zm.01 0v1h1v-1h-1Zm0-.01h1v-1h-1v1ZM10 11v-1H9v1h1Zm0 .01H9v1h1v-1Zm.01 0v1h1v-1h-1Zm0-.01h1v-1h-1v1ZM10 15v-1H9v1h1Zm0 .01H9v1h1v-1Zm.01 0v1h1v-1h-1Zm0-.01h1v-1h-1v1ZM15 15v-1h-1v1h1Zm0 .01h-1v1h1v-1Zm.01 0v1h1v-1h-1Zm0-.01h1v-1h-1v1ZM15 11v-1h-1v1h1Zm0 .01h-1v1h1v-1Zm.01 0v1h1v-1h-1Zm0-.01h1v-1h-1v1ZM5 15v-1H4v1h1Zm0 .01H4v1h1v-1Zm.01 0v1h1v-1h-1Zm0-.01h1v-1h-1v1ZM2 4h16V2H2v2Zm16 0h2a2 2 0 0 0-2-2v2Zm0 0v14h2V4h-2Zm0 14v2a2 2 0 0 0 2-2h-2Zm0 0H2v2h16v-2ZM2 18H0a2 2 0 0 0 2 2v-2Zm0 0V4H0v14h2ZM2 4V2a2 2 0 0 0-2 2h2Zm2-3v3h2V1H4Zm5 0v3h2V1H9Zm5 0v3h2V1h-2ZM1 8h18V6H1v2Zm3 3v.01h2V11H4Zm1 1.01h.01v-2H5v2Zm1.01-1V11h-2v.01h2Zm-1-1.01H5v2h.01v-2ZM9 11v.01h2V11H9Zm1 1.01h.01v-2H10v2Zm1.01-1V11h-2v.01h2Zm-1-1.01H10v2h.01v-2ZM9 15v.01h2V15H9Zm1 1.01h.01v-2H10v2Zm1.01-1V15h-2v.01h2Zm-1-1.01H10v2h.01v-2ZM14 15v.01h2V15h-2Zm1 1.01h.01v-2H15v2Zm1.01-1V15h-2v.01h2Zm-1-1.01H15v2h.01v-2ZM14 11v.01h2V11h-2Zm1 1.01h.01v-2H15v2Zm1.01-1V11h-2v.01h2Zm-1-1.01H15v2h.01v-2ZM4 15v.01h2V15H4Zm1 1.01h.01v-2H5v2Zm1.01-1V15h-2v.01h2Zm-1-1.01H5v2h.01v-2Z"
                />
              </svg>
            </div>
            <div className=" sm:text-left text-center mt-6 sm:mt-0">
              <h2 className="text-gray-900 text-center text-lg title-font font-medium mb-1">
                Date
              </h2>
              <p className="leading-relaxed text-base text-center">
                {format(new Date(event.Date), "dd/MM/yyyy HH:mm")}
              </p>
            </div>
          </div>
          <div className="flex justify-start  flex-1   flex-col items-center  border-b   border-gray-200  ">
            <div className="sm:w-32 sm:h-32 h-20 w-20  inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 flex-shrink-0">
              <svg
                fill="none"
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                className="sm:w-16 sm:h-16 w-10 h-10"
                viewBox="0 0 24 24"
              >
                <path d="M20 21v-2a4 4 0 00-4-4H8a4 4 0 00-4 4v2"></path>
                <circle cx="12" cy="7" r="4"></circle>
              </svg>
            </div>
            <div className=" sm:text-left text-center mt-6 sm:mt-0">
              <h2 className="text-gray-900 text-center text-lg title-font font-medium mb-1">
                The Organizer
              </h2>
              <p className="leading-relaxed text-base text-center">
                {event.Organizer.FirstName + " " + event.Organizer.LastName}
              </p>
            </div>
          </div>
          <div className="flex justify-start flex-1 flex-col items-center  border-b   border-gray-200  ">
            <div className="sm:w-32 sm:h-32 h-20 w-20  inline-flex items-center justify-center rounded-full bg-indigo-100 text-indigo-500 flex-shrink-0">
              <svg
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                className="sm:w-16 sm:h-16 w-10 h-10"
                fill="none"
                viewBox="0 0 30 30"
              >
                <path d="M11.163 11.554c0-2.655 2.166-4.807 4.837-4.807s4.837 2.152 4.837 4.807-2.166 4.806-4.837 4.806-4.837-2.152-4.837-4.806zM7.777 12.154c0 2.011 2.454 6.25 2.454 6.25l5.769 9.614 5.438-9.613c0 0 2.785-4.27 2.785-6.25 0-4.513-3.682-8.171-8.223-8.171s-8.223 3.657-8.223 8.17z"></path>
              </svg>
            </div>
            <div className=" sm:text-left text-center mt-6 sm:mt-0">
              <h2 className="text-gray-900 text-center text-lg title-font font-medium mb-1">
                Location
              </h2>
              <p className="leading-relaxed text-base text-center">
                {event.Place.City +
                  ", " +
                  event.Place.Country +
                  ", " +
                  event.Place.Address +
                  ", " +
                  event.Place.AddressNumber}
              </p>
            </div>
          </div>

          <div />
        </div>
        <div className="inline-flex w-full justify-center mt-5">
          {!event.State && (
            <a className="flex items-center  px-3 mr-1 py-2 bg-red-500 text-white font-semibold rounded">
              <svg
                fill="none"
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="3"
                className="w-5 h-5 mr-2"
                viewBox="0 0 24 24"
              >
                <path d="M20 6L9 17l-5-5"></path>
              </svg>
              In draft
            </a>
          )}

          {new Date(event.Date).getTime() < new Date().getTime() && (
            <>
              <a className="flex items-center  px-3 mr-1 py-2 bg-red-500 text-white font-semibold rounded">
                <svg
                  fill="none"
                  stroke="currentColor"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="3"
                  className="w-5 h-5 mr-2"
                  viewBox="0 0 24 24"
                >
                  <path d="M20 6L9 17l-5-5"></path>
                </svg>
                Event finished
              </a>
            </>
          )}

          <button
            onClick={() => {
              console.log("cambiando el estado")
              idEventEdit.setSubject(event.Id)
              alertLoginModalSubject.setSubject(true)
            }}
            className="flex px-3 py-2 bg-indigo-500 mr-1 text-white font-semibold rounded"
          >
            <svg
              className="w-6 h-6 text-gray-800 dark:text-white"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 21 21"
            >
              <path
                stroke="currentColor"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="1.5"
                d="M7.418 17.861 1 20l2.139-6.418m4.279 4.279 10.7-10.7a3.027 3.027 0 0 0-2.14-5.165c-.802 0-1.571.319-2.139.886l-10.7 10.7m4.279 4.279-4.279-4.279m2.139 2.14 7.844-7.844m-1.426-2.853 4.279 4.279"
              />
            </svg>
            <span className="ml-1">Edit</span>
          </button>
          <button
            onClick={() => showDetailsModalSubject.setSubject(false)}
            className="flex px-3 py-2 bg-red-500 mr-1 text-white font-semibold rounded"
          >
            <svg
              className="w-6 h-6 text-gray-800 dark:text-white"
              aria-hidden="true"
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 12 16"
            >
              <path
                stroke="currentColor"
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="1.5"
                d="M1 1v14m8.336-.479-6.5-5.774a1 1 0 0 1 0-1.494l6.5-5.774A1 1 0 0 1 11 2.227v11.546a1 1 0 0 1-1.664.748Z"
              />
            </svg>
            <span>Back</span>
          </button>
        </div>
      </div>
    </section>
  )
}

export default ModalDetailEventAdmin
