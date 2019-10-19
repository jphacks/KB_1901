import {NgModule} from '@angular/core';
import {PreloadAllModules, RouterModule, Routes} from '@angular/router';

const routes: Routes = [
    {path: '', redirectTo: 'home', pathMatch: 'full'},
    {path: 'home', loadChildren: () => import('./home/home.module').then(m => m.HomePageModule)},
    {path: 'login', loadChildren: './login/login.module#LoginPageModule'},
    {path: 'schedule-new', loadChildren: './schedule-new/schedule-new.module#ScheduleNewPageModule'},
    {path: 'schedule-list', loadChildren: './schedule-list/schedule-list.module#ScheduleListPageModule'},
    {path: 'schedule-id', loadChildren: './schedule-id/schedule-id.module#ScheduleIdPageModule'},
    {path: 'result', loadChildren: './result/result.module#ResultPageModule'},
    {path: 'found-store', loadChildren: './found-store/found-store.module#FoundStorePageModule'},
    {path: 'schedule-form', loadChildren: './schedule-form/schedule-form.module#ScheduleFormPageModule'},
];

@NgModule({
    imports: [
        RouterModule.forRoot(routes, {preloadingStrategy: PreloadAllModules})
    ],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
