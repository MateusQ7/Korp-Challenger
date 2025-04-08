import { Component, EventEmitter, Output, ViewChild, ElementRef, AfterViewInit, Input, OnChanges, SimpleChanges } from '@angular/core';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { CommonModule } from '@angular/common';
import * as bootstrap from 'bootstrap';
import { Product } from '../../models/product.model';
import { ProductService } from '../../pages/product/product.service';

@Component({
  selector: 'app-product-form',
  standalone: true,
  imports: [CommonModule, ReactiveFormsModule, FormsModule],
  templateUrl: './product-form.component.html',
  styleUrl: './product-form.component.css'
})
export class ProductFormComponent implements AfterViewInit, OnChanges {

  @ViewChild('productModal', { static: false }) modalElement!: ElementRef<HTMLDivElement>;
  @Output() productSaved = new EventEmitter<Product>();

  productForm: FormGroup;
  private modalInstance: bootstrap.Modal | null = null;

  constructor(
    private fb: FormBuilder,
    private productService: ProductService
  ) {
    this.productForm = this.fb.group({
      name: ['', Validators.required],
      price: [0, [Validators.required, Validators.min(0)]],
      stock: [0, [Validators.required, Validators.min(0)]],
    });
  }
  ngOnChanges(changes: SimpleChanges): void {
    throw new Error('Method not implemented.');
  }

  ngAfterViewInit(): void {
    if (this.modalElement?.nativeElement) {
      this.modalInstance = new bootstrap.Modal(this.modalElement.nativeElement);
    }
  }

  openModal(): void {
    if (this.modalInstance) {
      this.modalInstance.show();
    } else {
      console.error("Modal de produto não foi inicializado.");
    }
  }

  closeModal(): void {
    if (this.modalInstance) {
      this.modalInstance.hide();
    }
  }

  saveProduct(): void {
    if (this.productForm.valid) {
      const newProduct: Product = this.productForm.value;

      this.productService.create(newProduct).subscribe({
        next: (res: Product) => {
          console.log('Produto criado com sucesso:', res);
          this.productSaved.emit(res);
          setTimeout(() => {
            window.location.reload();
          }, 300);
        },
        error: (err) => {
          console.error('Erro ao criar produto:', err);
        },
        complete: () => {
          this.closeModal();
          this.productForm.reset();
        }
      });
    } else {
      console.warn('Formulário de produto inválido.');
    }
  }
}
