import { clienteAxios } from "../client.ts"
import { SubscribeAddDTO } from "../dtos/input.ts"
import { UserDTO } from "../dtos/output.ts"

export const getAllEventsPublished = () => {
  return clienteAxios.get("/api/v1/events").then((r) => r.data)
}

export const SubscribeToEvent = (data: SubscribeAddDTO) => {
  return clienteAxios.post("/api/v1/events", data).then((r) => r.data)
}

export const GetMyEvents = (state?: string) => {
  const id = (JSON.parse(localStorage.getItem("user") ?? "") as UserDTO).Id
  return clienteAxios
    .get(`/api/v1/events/user/${id}?state=${state}`)
    .then((r) => r.data)
}

export const GetEventsFiltered = (
    title: string,
  date: string,
  state: string,
) => {
  return clienteAxios
    .get(`/api/v1/events?title=${title??""}&date=${date??""}&state=${state??""}`)
    .then((r) => r.data)
}


export const GetEventsAdmin = () => {

  return clienteAxios
      .get("/api/v1/admin/events", {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': "Bearer " + localStorage.getItem("token"),
        },
      })
      .then((r) => r.data)
}


export const EditService =(id: number, data: any ) => {
  return clienteAxios
    .put("/api/v1/admin/events/"+id, data, {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': "Bearer " + localStorage.getItem("token"),
      },
    })
    .then((r) => r.data)
}


export const CreateEvent = (data: any) => {
  return clienteAxios
    .post("/api/v1/admin/events", data, {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': "Bearer " + localStorage.getItem("token"),
      },
    })
    .then((r) => r.data)
}