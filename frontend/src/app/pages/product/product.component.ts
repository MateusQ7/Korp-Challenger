import { Component, EventEmitter, Input, OnInit, Output, ViewChild } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Product } from '../../models/product.model';
import { ProductService } from './product.service';
import { ProductFormComponent } from '../../components/product-form/product-form.component';
import { InvoiceFormComponent } from '../../components/invoice-form/invoice-form.component';
import { Invoice } from '../../models/invoice.model';
import { InvoiceFormService } from '../../components/invoice-form/invoice-form.service';

@Component({
  selector: 'app-product',
  standalone: true,
  imports: [CommonModule, ProductFormComponent, InvoiceFormComponent],
  templateUrl: './product.component.html',
  styleUrl: './product.component.css'
})
export class ProductComponent implements OnInit {
  products: Product[] = [];
  productsMap: { [key: number]: Product } = {};
  invoices: Invoice[] = [];
  error: string | undefined = undefined;

  @ViewChild(ProductFormComponent) productModal!: ProductFormComponent;
  @ViewChild('invoiceModal') invoiceModal!: InvoiceFormComponent;

  constructor(
    private readonly productService: ProductService,
    private readonly invoiceFormService: InvoiceFormService,
  ) {}

  ngOnInit(): void {
    this.getProducts();
  }

  getProducts(): void {
    this.productService.getAll().subscribe({
      next: (data: Product[]) => {
        this.products = data;
        this.productsMap = {};
        data.forEach(p => this.productsMap[p.id!] = p);
        this.getAllInvoices();
      },
      error: (error) => {
        this.error = error;
        console.error('Erro ao buscar produtos:', error);
      },
      complete: () => {
        console.log('Carregamento de produtos concluído.');
      }
    });
  }

  getAllInvoices(): void {
    this.invoiceFormService.getAllInvoices().subscribe({
      next: (data: any[]) => {
        this.invoices = data.map(invoice => ({
          ...invoice,
          products: invoice.products.map((product: any) => ({
            product: this.productsMap[product.product_id],
            quantity: product.quantity
          }))
        }));
      },
      error: (error) => {
        this.error = error;
        console.error('Erro ao buscar invoices:', error);
      },
      complete: () => {
        console.log('Carregamento de invoices concluído.');
      }
    });
  }

  openProductModal(): void {
    this.productModal.openModal();
  }

  openInvoiceModal(): void {
    this.invoiceModal.openModal();
  }

  addProduct(product: Product): void {
    this.products.push(product);
    this.productsMap[product.id!] = product;
  }

  updateProductInList(updatedProduct: Product): void {
    const index = this.products.findIndex(p => p.id === updatedProduct.id);
    if (index !== -1) {
      this.products[index] = updatedProduct;
      this.productsMap[updatedProduct.id!] = updatedProduct;
    }
  }

  onInvoiceSaved(invoice: Invoice): void {
    invoice.products = invoice.products.map(product => ({
      product: this.productsMap[product.product.id!],
      quantity: product.quantity
    }));
    this.invoices.push(invoice);
  }

  getInvoiceTotal(invoice: Invoice): number {
    return invoice.products?.reduce((total, product) => total + (product.product.price * product.quantity), 0) || 0;
  }

  printInvoice(invoiceId?: number): void {
    if (!invoiceId) {
      console.error('ID da nota fiscal é inválido.');
      return;
    }

    this.invoiceFormService.printInvoice(invoiceId).subscribe({
        next: () => {
          console.log('Nota fiscal enviada para impressão!');
          setTimeout(() => {
            window.location.reload();
          }, 300);
        },
        error: (err) => {
          console.error('Erro ao imprimir nota fiscal:', err);
        }
      });
  }
}
