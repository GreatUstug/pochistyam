export type Room = {
    id: number
    name: string
    width: number
    height: number
    dirtiness: number
    needsCleaning: boolean
}

export type RoomRequest = {
    clientId: string
    roomId: number
}