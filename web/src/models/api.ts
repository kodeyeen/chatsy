export interface Page<T> {
    items: T[]
    count: number
    limit: number
    offset: number
}
