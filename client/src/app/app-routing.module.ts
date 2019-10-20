import {NgModule} from '@angular/core';
import {PreloadAllModules, RouterModule, Routes} from '@angular/router';

const routes: Routes = [
    {path: '', redirectTo: 'home/null/null', pathMatch: 'full'},
    {path: 'home/:auth_token/:user_name', loadChildren: () => import('./home/home.module').then(m => m.HomePageModule)},
    {path: 'login', loadChildren: './login/login.module#LoginPageModule'},
    {path: 'schedule-new/:auth_token/:user_name', loadChildren: './schedule-new/schedule-new.module#ScheduleNewPageModule'},
    {path: 'schedule-list/:auth_token/:user_name', loadChildren: './schedule-list/schedule-list.module#ScheduleListPageModule'},
    {path: 'schedule-id/:auth_token/:key', loadChildren: './schedule-id/schedule-id.module#ScheduleIdPageModule'},
    {path: 'result/:json_data', loadChildren: './result/result.module#ResultPageModule'},
    {path: 'found-store/:area/:genre', loadChildren: './found-store/found-store.module#FoundStorePageModule'},
    {path: 'schedule-form/:key', loadChildren: './schedule-form/schedule-form.module#ScheduleFormPageModule'},
];

@NgModule({
    imports: [
        RouterModule.forRoot(routes, {preloadingStrategy: PreloadAllModules})
    ],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
