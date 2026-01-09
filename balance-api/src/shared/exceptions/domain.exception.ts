/** Domain exception base class */
export class DomainException extends Error {
  constructor(message: string) {
    super(message);
  }
}