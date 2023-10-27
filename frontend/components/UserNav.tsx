import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Button } from '@/components/ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu';
import { signOut, useSession } from 'next-auth/react';
import { ReactElement } from 'react';

export function UserNav(): ReactElement {
  const session = useSession();

  const avatarImageSrc = session.data?.user?.image ?? '';
  const avatarImageAlt = session.data?.user?.name ?? 'unknown';
  const avatarFallback = getAvatarFallback(session.data?.user?.name);

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button variant="ghost" className="relative h-8 w-8 rounded-full">
          <Avatar className="h-8 w-8">
            <AvatarImage src={avatarImageSrc} alt={avatarImageAlt} />
            <AvatarFallback>{avatarFallback}</AvatarFallback>
          </Avatar>
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className="w-56" align="end" forceMount>
        <DropdownMenuLabel className="font-normal">
          <div className="flex flex-col space-y-1">
            {session.data?.user?.name && (
              <p className="text-sm font-medium leading-none">
                {session.data.user.name}
              </p>
            )}
            {session.data?.user?.email && (
              <p className="text-xs leading-none text-muted-foreground">
                {session.data.user.email}
              </p>
            )}
          </div>
        </DropdownMenuLabel>
        <DropdownMenuSeparator />
        <DropdownMenuGroup>
          <DropdownMenuItem disabled>Profile</DropdownMenuItem>
          <DropdownMenuItem disabled>Settings</DropdownMenuItem>
        </DropdownMenuGroup>
        <DropdownMenuSeparator />
        {session.status === 'authenticated' && (
          <DropdownMenuItem onClick={() => signOut()}>Log out</DropdownMenuItem>
        )}
      </DropdownMenuContent>
    </DropdownMenu>
  );
}

function getAvatarFallback(name?: string | null): string {
  return !!name ? name[0].toUpperCase() : 'UN'; // unknown?
}
