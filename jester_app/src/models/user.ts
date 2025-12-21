export type UserStatus = {
    id: string,
    name: string,
    description: string
}

export type User = {
    id: string,
    status_id: string,
    username: string,
    email: string,
}