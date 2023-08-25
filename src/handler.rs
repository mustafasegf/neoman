use crate::app::{App, AppResult, Selected};
use crossterm::event::{KeyCode, KeyEvent, KeyModifiers};

/// Handles the key events and updates the state of [`App`].
pub fn handle_key_events(key_event: KeyEvent, app: &mut App) -> AppResult<()> {
    match key_event.code {
        KeyCode::Esc | KeyCode::Char('q') => {
            app.quit();
        }
        KeyCode::Char('c') | KeyCode::Char('C') => {
            if key_event.modifiers == KeyModifiers::CONTROL {
                app.quit();
            }
        }

        KeyCode::Char(' ') | KeyCode::Char('o') | KeyCode::Enter => app.sidebar.tree.toggle(),

        KeyCode::Left => app.sidebar.tree.left(),
        KeyCode::Right => app.sidebar.tree.right(),
        KeyCode::Down => app.sidebar.tree.down(),
        KeyCode::Up => app.sidebar.tree.up(),
        KeyCode::Home => app.sidebar.tree.first(),
        KeyCode::End => app.sidebar.tree.last(),

        KeyCode::Char('b') | KeyCode::Char('B') => {
            if key_event.modifiers == KeyModifiers::CONTROL {
                app.toggle_sidebar();
            }
        }

        KeyCode::Tab => {
            app.selected = match app.selected {
                Selected::Sidebar => Selected::Tabs,
                Selected::Tabs => Selected::MethodBar,
                Selected::MethodBar => Selected::Urlbar,
                Selected::Urlbar => Selected::Requestbar,
                Selected::Requestbar => Selected::Sidebar,
            };
        }

        // Other handlers you could add here.
        _ => {}
    }
    Ok(())
}
