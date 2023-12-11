// Importa la biblioteca Axios
import axios from 'axios';

export const clienteAxios = axios.create({
    baseURL: import.meta.env.VITE_BASE_URL, // URL base para las solicitudes
    timeout: 15000, // Tiempo máximo de espera para la solicitud en milisegundos
    headers: {
        'Content-Type': 'application/json',
    },
});


