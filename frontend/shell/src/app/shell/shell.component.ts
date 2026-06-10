import { Component, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { SidebarComponent } from './sidebar/sidebar.component';

@Component({
  selector: 'app-shell',
  standalone: true,
  imports: [RouterOutlet, SidebarComponent],
  templateUrl: './shell.component.html',
})
export class ShellComponent {
  collapsed = signal(false);

  toggleSidebar() {
    this.collapsed.update((v) => !v);
  }
}