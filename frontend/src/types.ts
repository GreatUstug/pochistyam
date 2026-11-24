export type Room = {
    id: string;
    name: string;
    width: number;
    height: number;
    dirtiness: number;
    needsCleaning: boolean;
};

export type CreateRoomRequest = {
    name: string;
    width: number;
    height: number;
};

export type RoomsListResponse = {
    rooms: Room[];
};