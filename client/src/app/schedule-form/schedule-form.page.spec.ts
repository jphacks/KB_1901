import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ScheduleFormPage } from './schedule-form.page';

describe('ScheduleFormPage', () => {
  let component: ScheduleFormPage;
  let fixture: ComponentFixture<ScheduleFormPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ScheduleFormPage ],
      schemas: [CUSTOM_ELEMENTS_SCHEMA],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ScheduleFormPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
