
export interface EventDTO {
    Id: number;
    Title: string;
    ShortDescription: string;
    LongDescription: string;
    Date: Date;
    Organizer: OrganizerDTO;
    Place: PlaceDTO;
    State: boolean;
}
export interface PlaceDTO {
    Id: number;
    Address: string;
    AddressNumber: number;
    City: string;
    Country: string;
}

export interface OrganizerDTO {
    Id: number;
    FirstName: string;
    LastName: string;
}

export  interface UserDTO {
    Id: number;
    FirstName: string;
    LastName: string;
    Email: string;
    Role: string;
    Events: EventDTO[];
}

export interface AuthResponse {
    User: UserDTO;
    Token: LoginResponse;
}

export interface LoginResponse {
    AccessToken: string;
    RefreshToken: string;
}

