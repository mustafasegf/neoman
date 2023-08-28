use crate::app::{App, AppResult, Selected};
use crossterm::event::{KeyCode, KeyEvent, KeyModifiers};

/// Handles the key events and updates the state of [`App`].
pub fn handle_key_events(key_event: KeyEvent, app: &mut App) -> AppResult<()> {
    // global key handlers
    match key_event.code {
        KeyCode::Char('c') | KeyCode::Char('C') => {
            if key_event.modifiers == KeyModifiers::CONTROL {
                app.quit();
            }
        }

        KeyCode::Char('b') | KeyCode::Char('B') => {
            if key_event.modifiers == KeyModifiers::CONTROL {
                app.toggle_sidebar();
            }
        }

        KeyCode::Esc | KeyCode::Char('q') => {
            app.quit();
        }

        KeyCode::Tab => {
            if key_event.modifiers == KeyModifiers::CONTROL {
                app.selected = match app.selected {
                    Selected::Sidebar => Selected::Responsebar,
                    Selected::Tabs => Selected::Sidebar,
                    Selected::MethodBar => Selected::Tabs,
                    Selected::Urlbar => Selected::MethodBar,
                    Selected::Requestbar => Selected::Urlbar,
                    Selected::Responsebar => Selected::Requestbar,
                };
            } else {
                app.selected = match app.selected {
                    Selected::Sidebar => Selected::Tabs,
                    Selected::Tabs => Selected::MethodBar,
                    Selected::MethodBar => Selected::Urlbar,
                    Selected::Urlbar => Selected::Requestbar,
                    Selected::Requestbar => Selected::Responsebar,
                    Selected::Responsebar => Selected::Sidebar,
                };
            }
        }

        // Other handlers you could add here.
        _ => {}
    }

    match app.selected {
        Selected::Sidebar => match key_event.code {
            KeyCode::Char(' ') | KeyCode::Char('o') | KeyCode::Enter => {
                if let Some(item) = app.sidebar.selected() {
                    if item.children().is_empty() {
                        // app.tabs.add(item.inner().clone());
                        let item_name = item.inner().to_string();
                        match app
                            .tabs
                            .tabs
                            .iter()
                            .enumerate()
                            .find(|(_, item)| item.borrow().name == item_name)
                            .map(|(i, _)| i)
                        {
                            Some(i) => {
                                app.tabs.selected = i;
                            }
                            None => {
                                app.tabs.add(item.inner().clone());
                                app.tabs.selected = app.tabs.tabs.len() - 1;
                                app.selected = Selected::Tabs;
                            }
                        }
                    } else {
                        app.sidebar.tree.toggle();
                    }
                }
            }

            KeyCode::Left => app.sidebar.tree.left(),
            KeyCode::Right => app.sidebar.tree.right(),
            KeyCode::Down => app.sidebar.tree.down(),
            KeyCode::Up => app.sidebar.tree.up(),
            KeyCode::Home => app.sidebar.tree.first(),
            KeyCode::End => app.sidebar.tree.last(),

            _ => {}
        },
        Selected::Tabs => match key_event.code {
            KeyCode::Left => {
                app.tabs.left();
            }
            KeyCode::Right => {
                app.tabs.right();
            }
            KeyCode::Down => {
                app.tabs.right();
            }
            KeyCode::Up => {
                app.tabs.left();
            }
            KeyCode::Home => {
                app.tabs.first();
            }
            KeyCode::End => {
                app.tabs.last();
            }
            KeyCode::Enter => {}
            _ => {}
        },
        Selected::MethodBar => {}
        Selected::Urlbar => {}
        Selected::Requestbar => {}
        Selected::Responsebar => {}
    }
    Ok(())
}
