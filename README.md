# D2 ModMaker

The idea behind this project is to play Diablo II the way that you want to play.  Every option in the `ModConfig` is optional.

# Links
* [Repository](https://github.com/tlentz/d2modmaker)
* [Releases](https://github.com/tlentz/d2modmaker/releases)
* [Installation Guide](https://docs.google.com/document/d/1M5uY67giX4DGnXHxmApb-Uf5AUZdN5yquidAr2BUR_c/edit?usp=sharing)
* [Discord Community](https://discord.gg/fePUH3)

# ModConfig

The mod config is located in `cfg.json`.  You can change this config to your liking to produce a new `data` folder.

## ModConfig Options
* **IncreaseStackSizes** `bool`
    * Increases book of tp to 100
    * Increases book of id to 100
    * Increases arrows maxstack to 511
    * Increases bolts maxstack to 511
    * Increases key stack sizes to 100
* **IncreaseMonsterDensity** `float`
    * Will increase the density of all areas by the given multiplier
    * `MAX: 30.0`
    * `MIN: 0.0`
    * Set to `-1` to omit
* **EnableTownSkills** `bool`
    * Enables all skills in town
* **NoDropZero** `bool`
    * Sets "NoDrop" = 0 (Monsters will always drop items)
* **QuestDrops** `bool`
    * Enables quest drops for boss kills always
* **UniqueItemDropRate** `float`
    * Will increase the rate in which uniques/sets drop
    * Set to `-1` to omit
* **StartWithCube** `bool`
    * Characters will start with cube when created
* **RandomOptions** `RandomOptions`
    * **Randomize** `bool`
        * Will randomize if set to true
    * **Seed** `int`
        * Will use this seed for randomization
        * Set to `-1` to generate a random seed
    * **IsBalanced** `bool`
        * bucketizes props by levels `[0-30] [31-60] [61+]` so that you don't get crazy hell stats on normal items, but still get a wide range of randomization
    * **MinProps** `int`
        * Minimum number of non blank props that spawn on an item
        * Set to `-1` to omit
    * **MaxProps** `int`
        * Maximum number of non blank props that spawn on an item
        * Set to `-1` to omit
    * **UseOSkills** `bool`
        * Will change class only skills to oskills
    * **PerfectProps** `bool`
        * All props will have the max value for min/max values

## Example ModConfig
```json
 {
    "IncreaseStackSizes": true,
    "IncreaseMonsterDensity": 1,
    "EnableTownSkills": true,
    "NoDropZero": true,
    "QuestDrops": true,
    "UniqueItemDropRate": 100,
    "StartWithCube": true,
    "RandomOptions": {
        "Randomize": true,
        "Seed": -1,
        "IsBalanced": false,
        "MinProps": -1,
        "MaxProps": -1,
        "UseOSkills": true,
        "PerfectProps": false
    }
}
```

# Screenshots
### Nagel
![Nagel](https://i.imgur.com/1zOKK3q.png)
### Raven Claw
![Raven Claw](https://i.imgur.com/tmxZpjc.png)
### Venom Ward
![Venom Ward](https://i.imgur.com/7cLQDBN.png)
### Angelic Halo
![Angelic Halo](https://i.imgur.com/N3Om8II.png)
### Wall of Eyeless
![Wall of Eyeless](https://i.imgur.com/QL07TKL.png)
### MonsterDensity: 30
![MonsterDensity: 30](https://i.imgur.com/d6iCBZA.png)
