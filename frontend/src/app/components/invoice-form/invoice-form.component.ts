import { Component, EventEmitter, Output, ViewChild, ElementRef, AfterViewInit, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import * as bootstrap from 'bootstrap';
import { Product } from '../../models/product.model';
import { Invoice } from '../../models/invoice.model';
import { ProductService } from '../../pages/product/product.service';
import { InvoiceFormService } from './invoice-form.service';
import { InvoiceRequest } from '../../models/invoiceRequest.model';

@Component({
  selector: 'app-invoice-form',
  standalone: true,
  imports: [CommonModule, FormsModule],
  templateUrl: './invoice-form.component.html',
  styleUrl: './invoice-form.component.css'
})
export class InvoiceFormComponent implements AfterViewInit {
  @ViewChild('invoiceModal', { static: false }) modalElement!: ElementRef<HTMLDivElement>;
  @Input() products: Product[] = [];
  @Output() invoiceSaved = new EventEmitter<Invoice>();

  selectedProducts: { [productId: number]: number } = {};
  private modalInstance: bootstrap.Modal | null = null;

  constructor(
    private productService: ProductService,
    private invoiceFormService: InvoiceFormService
  ) {}

  ngAfterViewInit(): void {
    if (this.modalElement?.nativeElement) {
      this.modalInstance = new bootstrap.Modal(this.modalElement.nativeElement);
    }
  }

  openModal(): void {
    this.loadProducts();
    if (this.modalInstance) this.modalInstance.show();
  }

  closeModal(): void {
    if (this.modalInstance) this.modalInstance.hide();
  }

  loadProducts(): void {
    this.productService.getAll().subscribe({
      next: (res) => (this.products = res),
      error: (err) => console.error('Erro ao carregar produtos:', err),
    });
  }

  toggleProductSelection(productId: number, isChecked: boolean): void {
    if (isChecked) {
      this.selectedProducts[productId] = 1;
    } else {
      delete this.selectedProducts[productId];
    }
  }

  updateQuantity(productId: number, quantity: number): void {
    if (this.selectedProducts[productId]) {
      this.selectedProducts[productId] = quantity;
    }
  }

  saveInvoice(): void {
    const selected = this.products.filter(p => this.selectedProducts.hasOwnProperty(p.id!));

    const invoiceRequest: InvoiceRequest = {
      number: 'NF-' + Date.now() + '-' + Math.floor(Math.random() * 1000),
      status: 'aberta',
      products: selected.map(p => ({
        product_id: p.id!,
        quantity: Number(this.selectedProducts[p.id!])
      }))
    };

    console.log(invoiceRequest);

    this.invoiceFormService.createInvoice(invoiceRequest).subscribe({
      next: (res) => {
        console.log('Nota fiscal criada com sucesso:', res);
        this.invoiceSaved.emit(res);
        this.closeModal();

        setTimeout(() => {
          window.location.reload();
        }, 300);
      },
      error: (err) => {
        console.error('Erro ao criar nota fiscal:', err);
      }
    });
  }

  getCheckboxValue(event: Event): boolean {
    return (event.target as HTMLInputElement).checked;
  }
}
