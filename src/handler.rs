use crate::app::{App, AppResult};
use crossterm::event::{KeyCode, KeyEvent, KeyModifiers};

/// Handles the key events and updates the state of [`App`].
pub fn handle_key_events(key_event: KeyEvent, app: &mut App) -> AppResult<()> {
    match key_event.code {
        // Exit application on `ESC` or `q`
        KeyCode::Esc | KeyCode::Char('q') => {
            app.quit();
        }
        // Exit application on `Ctrl-C`
        KeyCode::Char('c') | KeyCode::Char('C') => {
            if key_event.modifiers == KeyModifiers::CONTROL {
                app.quit();
            }
        }
        // Counter handlers
        KeyCode::Right => {
            app.increment_counter();
        }
        KeyCode::Left => {
            app.decrement_counter();
        }
        // Exit application on `Ctrl-C`
        KeyCode::Char('b') | KeyCode::Char('B') => {
            if key_event.modifiers == KeyModifiers::CONTROL {
                app.toggle_sidebar();
            }
        }

        // Other handlers you could add here.
        _ => {}
    }
    Ok(())
}
