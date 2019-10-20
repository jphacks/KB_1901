import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FoundStorePage } from './found-store.page';

describe('FoundStorePage', () => {
  let component: FoundStorePage;
  let fixture: ComponentFixture<FoundStorePage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FoundStorePage ],
      schemas: [CUSTOM_ELEMENTS_SCHEMA],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FoundStorePage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
