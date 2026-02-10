/**
 * A container of integers that should support
 * addition, removal, and search for the median integer
 */
class Container {
  private list: number[];
  
  constructor() {
    this.list = [];
  }

  /**
   * Adds the specified value to the container
   */
  add(value: number): void {
    for (let i = 0; i < this.list.length; i++) {
      if (this.list[i] > value) {
        this.list.splice(i, 0, value);
        return;
      }
    }

    this.list.push(value);
  }

  /**
   * Attempts to delete one item of the specified value from the container
   *
   * @return {boolean} true, if the value has been deleted, or
   *                   false, otherwise.
   */
  delete(value: number): boolean {
    for (let i = 0; i < this.list.length; i++) {
      if (this.list[i] === value) {
        this.list.splice(i, 1);
        return true;
      }
    }
    return false;
  }

  /**
   * Finds the container's median integer value, which is
   * the middle integer when the all integers are sorted in order.
   * If the sorted array has an even length,
   * the leftmost integer between the two middle
   * integers should be considered as the median.
   *
   * @return {number} the median if the array is not empty, or
   * @throws {Error} a runtime exception, otherwise.
   */
  getMedian(): number {
    var even = this.list.length % 2 === 0;
    var middlePosition = Math.floor(this.list.length / 2);
    
    if (this.list.length === 0) {
      throw new Error('Container is empty');
    }

    if (even) {
      return this.list[middlePosition - 1];
    }

    return this.list[middlePosition];
  }
}