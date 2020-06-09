export interface IhttpAPI {
  auth: IhttpAPIAuth
  events: IhttpAPIEvents
  sponsors: IhttpAPISponsors
  tickets: IhttpAPITickets
  users: IhttpAPIUsers
}

/**
 * API Promises
 *
 */

interface IhttpAPIAuth {
  login(username: string, password: string): Promise<IResponse>
  logout(): Promise<IResponse>
  profile(): Promise<IResponse<IUserInstance>>
  register(newUser: IUser): Promise<IResponse>
}

interface IhttpAPIEvents {
  add(newEvent: IEvent): Promise<IResponse>
  edit(eventID: string, update: IEvent): Promise<IResponse>
  get(eventID: string): Promise<IResponse<IEvent>>
  getBulk(limit: number, page: number): Promise<IResponse<IEventInstance[]>>
  remove(eventID: string): Promise<IResponse>
}

interface IhttpAPISponsors {
  add(newSponsor: ISponsor): Promise<IResponse>
  edit(sponsorID: string, update: IEvent): Promise<IResponse>
  get(sponsorID: string): Promise<IResponse<ISponsorInstance>>
  getBulk(limit: number, page: number): Promise<IResponse<ISponsorInstance[]>>
  remove(sponsorID: string): Promise<IResponse>
}
interface IhttpAPITickets {
  contact(newTicket: ITicket): Promise<IResponse>
  get(ticketID: string): Promise<IResponse<ITicketInstance>>
  response(
    ticketID: string,
    managerID: string,
    message: string,
  ): Promise<IResponse>
}

interface IhttpAPIUsers {
  edit(userID: string, update: IEvent): Promise<IResponse>
  get(userID: string): Promise<IResponse<IUserInstance>>
  getBulk(limit: number, page: number): Promise<IResponse<IUserInstance[]>>
  remove(userID: string): Promise<IResponse>
}

/**
 * API RESPONSES
 *
 */

interface IResponse<T = void> {
  data: T | IHTTPError
}

interface IHTTPError {
  title: string
  details: string
  invalidParams?: string
}

/**
 * MODELS
 *
 */

interface IUser {
  username: string
  password: string
  email: string
  fullname: string
}

interface ISponsor {
  name: string
  phones: string[]
  emails: string[]
}

interface IEvent {
  name: string
  description: string
}

interface ITicket {
  from: string
  fullname: string
  subject: string
  message: string
  responses: string[]
}

/**
 * INSTANCES
 *
 */

interface IUserInstance extends IUser {
  ID: string
}

interface IEventInstance extends IEvent {
  ID: string
}

interface ISponsorInstance extends ISponsor {
  ID: string
}

interface ITicketInstance extends ITicket {
  ID: string
}
