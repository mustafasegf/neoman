use crate::app::{App, AppResult};
use crossterm::event::{KeyCode, KeyEvent, KeyEventKind, KeyModifiers};

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
            if key_event.modifiers == KeyModifiers::CONTROL && key_event.kind == KeyEventKind::Press
            {
                app.toggle_sidebar();
            }
        }

        KeyCode::Down if key_event.kind == KeyEventKind::Press => {
            app.sidebar
                .list
                .get_current()
                .value()
                .unwrap()
                .borrow_mut()
                .selected = false;

            app.sidebar
                .list
                .move_next_and_get_always()
                .unwrap()
                .borrow_mut()
                .selected = true;
        }

        KeyCode::Up if key_event.kind == KeyEventKind::Press => {
          app.sidebar
              .list
              .get_current()
              .value()
              .unwrap()
              .borrow_mut()
              .selected = false;

          app.sidebar
              .list
              .move_prev_and_get_always()
              .unwrap()
              .borrow_mut()
              .selected = true;
      }


        // Other handlers you could add here.
        _ => {}
    }
    Ok(())
}
