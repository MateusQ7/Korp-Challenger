import { Product } from "./product.model";

export interface Invoice {
  id?: number;
  number: string;
  status: 'aberta' | 'fechada';
  date?: Date;
  total?: number;
  products: {
    product: Product;
    quantity: number;
  }[];
}
