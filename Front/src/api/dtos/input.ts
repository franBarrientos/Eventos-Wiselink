export interface LoginDTO {
    Email: string;
    Password: string;
}

export interface UserAddDTO {
    FirstName: string;
    LastName: string;
    Email: string;
    Password: string;
}

export interface SubscribeAddDTO {
    User: number;
    Event: number;
}

export interface PlaceAddDTO {
    Address: string;
    AddressNumber: number;
    City: string;
    Country: string;
}

export interface OrganizerAddDTO {
    FirstName: string;
    LastName: string;
}

export interface EventAddDTO {
    Title: string;
    ShortDescription: string;
    LongDescription: string;
    Date: string;
    Organizer: OrganizerAddDTO;
    Place: PlaceAddDTO;
    State: boolean;
}