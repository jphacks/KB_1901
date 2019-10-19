import { NgModule } from '@angular/core';
import { PreloadAllModules, RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  { path: '', redirectTo: 'login', pathMatch: 'full' },
  { path: 'login', loadChildren: () => import('./login/login.module').then( m => m.LoginPageModule)},
  { path: 'home', loadChildren: './home/home.module#HomePageModule' },
  { path: 'schedule-new', loadChildren: './schedule-new/schedule-new.module#ScheduleNewPageModule' },
  { path: 'schedule-list', loadChildren: './schedule-list/schedule-list.module#ScheduleListPageModule' },
  { path: 'schedule-id', loadChildren: './schedule-id/schedule-id.module#ScheduleIdPageModule' },
  { path: 'result', loadChildren: './result/result.module#ResultPageModule' },
  { path: 'found-store', loadChildren: './found-store/found-store.module#FoundStorePageModule' },
];

@NgModule({
  imports: [
    RouterModule.forRoot(routes, { preloadingStrategy: PreloadAllModules })
  ],
  exports: [RouterModule]
})
export class AppRoutingModule { }
