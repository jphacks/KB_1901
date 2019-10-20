import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ScheduleNewPage } from './schedule-new.page';

describe('ScheduleNewPage', () => {
  let component: ScheduleNewPage;
  let fixture: ComponentFixture<ScheduleNewPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ScheduleNewPage ],
      schemas: [CUSTOM_ELEMENTS_SCHEMA],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ScheduleNewPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
