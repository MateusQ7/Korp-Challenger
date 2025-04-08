import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Invoice } from '../../models/invoice.model';
import { Observable } from 'rxjs';
import { InvoiceRequest } from '../../models/invoiceRequest.model';

@Injectable({
  providedIn: 'root'
})
export class InvoiceFormService {

  private baseUrl = 'http://localhost:8082/invoices';

  constructor(
    private http: HttpClient
  ) { }

  createInvoice(invoiceR: InvoiceRequest): Observable<Invoice> {
    return this.http.post<Invoice>(this.baseUrl, invoiceR);
  }

  getInvoiceById(id: number): Observable<Invoice> {
    return this.http.get<Invoice>(`${this.baseUrl}/${id}`);
  }

  getAllInvoices(): Observable<Invoice[]> {
    return this.http.get<Invoice[]>(`${this.baseUrl}`)
  }

  printInvoice(id: number): Observable<Invoice> {
    return this.http.put<Invoice>(`${this.baseUrl}/${id}/print`, {});
  }

}
