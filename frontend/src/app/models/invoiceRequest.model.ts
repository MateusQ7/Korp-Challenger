export interface InvoiceRequest {
  number: string;
  status: string;
  products: {
    product_id: number;
    quantity: number;
  }[];
}
