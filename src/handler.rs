use crate::{
    app::{App, AppResult, Selected},
    component::urlbar::InputMode,
};
use crossterm::event::{KeyCode, KeyEvent, KeyModifiers};

/// Handles the key events and updates the state of [`App`].
pub async fn handle_key_events(key_event: KeyEvent, app: &mut App) -> AppResult<()> {
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
            if !app.urlbar.method_menu.is_open() {
                app.quit();
            }
        }

        KeyCode::Tab | KeyCode::Char('.') | KeyCode::Char(']') => {
            app.selected = match app.selected {
                Selected::Sidebar => Selected::Tabs,
                Selected::Tabs => Selected::MethodBar,
                Selected::MethodBar => Selected::Urlbar,
                Selected::Urlbar => Selected::RequestTab,
                Selected::RequestTab => Selected::Requestbar,
                Selected::Requestbar => Selected::Responsebar,
                Selected::Responsebar => Selected::Sidebar,
            };
        }

        KeyCode::Char(',') | KeyCode::Char('[') => {
            app.selected = match app.selected {
                Selected::Sidebar => Selected::Responsebar,
                Selected::Tabs => Selected::Sidebar,
                Selected::MethodBar => Selected::Tabs,
                Selected::Urlbar => Selected::MethodBar,
                Selected::RequestTab => Selected::Urlbar,
                Selected::Requestbar => Selected::RequestTab,
                Selected::Responsebar => Selected::Requestbar,
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
                                app.selected = Selected::Tabs;
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
        Selected::MethodBar => {
            match key_event.code {
                KeyCode::Char('h') | KeyCode::Left => app.urlbar.method_menu.left(),
                KeyCode::Char('l') | KeyCode::Right => app.urlbar.method_menu.right(),
                KeyCode::Char('j') | KeyCode::Down => app.urlbar.method_menu.down(),
                KeyCode::Char('k') | KeyCode::Up => app.urlbar.method_menu.up(),
                KeyCode::Esc => app.urlbar.method_menu.reset(),
                KeyCode::Enter => app.urlbar.method_menu.select(),
                _ => {}
            };

            for e in app.urlbar.method_menu.drain_events() {
                match e {
                    tui_menu::MenuEvent::Selected(item) => {
                        app.urlbar.method_menu.set_child_name(0, item.to_string());
                        app.urlbar.method_menu.close();
                        app.urlbar.method = item;
                    }
                }
            }
        }
        Selected::Urlbar => match app.urlbar.input_mode {
            InputMode::Normal => match key_event.code {
                KeyCode::Enter | KeyCode::Char('i') => app.urlbar.input_mode = InputMode::Insert,
                KeyCode::Char('o') => {
                    app.request().await;
                }
                _ => {}
            },
            InputMode::Insert => {
                match key_event.code {
                    // KeyCode::Esc => app.urlbar.input_mode = InputMode::Normal,
                    KeyCode::Enter => app.urlbar.input_mode = InputMode::Normal,
                    KeyCode::Char(c) => {
                        app.urlbar.text.insert(app.urlbar.cursor_position, c);
                        app.urlbar.cursor_position += 1;
                    }
                    KeyCode::Backspace => {
                        if app.urlbar.cursor_position > 0 {
                            app.urlbar.cursor_position -= 1;
                            app.urlbar.text.remove(app.urlbar.cursor_position);
                        }
                    }
                    KeyCode::Delete => {
                        if app.urlbar.cursor_position < app.urlbar.text.len() {
                            app.urlbar.text.remove(app.urlbar.cursor_position);
                        }
                    }
                    KeyCode::Left => {
                        if app.urlbar.cursor_position > 0 {
                            app.urlbar.cursor_position -= 1;
                        }
                    }
                    KeyCode::Right => {
                        if app.urlbar.cursor_position < app.urlbar.text.len() {
                            app.urlbar.cursor_position += 1;
                        }
                    }
                    KeyCode::Home => {
                        app.urlbar.cursor_position = 0;
                    }
                    KeyCode::End => {
                        app.urlbar.cursor_position = app.urlbar.text.len();
                    }
                    _ => {}
                }
            }
        },
        Selected::RequestTab => {
            match key_event.code {
                KeyCode::Char('h') | KeyCode::Left => app.requestbar.left(),
                KeyCode::Char('l') | KeyCode::Right => app.requestbar.right(),
                KeyCode::Char('j') | KeyCode::Down => app.requestbar.left(),
                KeyCode::Char('k') | KeyCode::Up => app.requestbar.right(),

                _ => {}
            };
        }
        Selected::Requestbar => {}
        Selected::Responsebar => {}
    }
    Ok(())
}
