import { Product } from "./product.model";

export interface InvoiceItem {
  product: Product;
  quantity: number;
}
