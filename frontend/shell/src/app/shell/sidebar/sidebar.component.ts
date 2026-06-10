import { Component, input, output } from '@angular/core';
import { RouterLink, RouterLinkActive } from '@angular/router';

interface NavItem {
  path: string;
  icon: string;
  label: string;
}

@Component({
  selector: 'app-sidebar',
  standalone: true,
  imports: [RouterLink, RouterLinkActive],
  templateUrl: './sidebar.component.html',
})
export class SidebarComponent {
  collapsed = input<boolean>(false);
  toggle = output<void>();

  navItems: NavItem[] = [
    { path: '/notes', icon: '🥞', label: 'Notes' },
    { path: '/scim',  icon: '👤', label: 'SCIM' },
  ];
}