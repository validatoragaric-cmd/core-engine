# main.py
from core_engine.config import Config
from core_engine.log import Log
from core_engine.utils import load_module

def main():
    Log.log_info("Core Engine init")

    config = Config()
    log_level = config.get("log_level")

    Log.log_info(f"Log level: {log_level}")

    modules = config.get("modules")
    for module in modules:
        Log.log_info(f"Loading module: {module}")
        module_path, module_name = module.split("/")
        module_path = f"core_engine.modules.{module_path}"
        module = load_module(module_path, module_name)
        module.init()

if __name__ == "__main__":
    main()